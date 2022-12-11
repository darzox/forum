package handlers

import (
	"html/template"
	"net/http"
	"strings"

	"forum/internal/model"
)

type Leaving interface {
	DeleteSession(userId uint) error
	UserBySession(cookie string) (*model.User, error)
}

type Logout struct {
	service Leaving
}

func CreateLogoutHandler(service Leaving) *Logout {
	return &Logout{
		service: service,
	}
}

func (l Logout) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorPage(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed, w)
		return
	}
	cookieFromClient := r.Header.Get("Cookie")
	cookieFromClient = strings.ReplaceAll(cookieFromClient, "Session-token=", "")
	user, err := l.service.UserBySession(cookieFromClient)
	if err != nil {
		errorPage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, w)
		return
	}
	err = l.service.DeleteSession(user.ID)
	if err != nil {
		errorPage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, w)
		return
	}
	t, err := template.New("index.html").Funcs(template.FuncMap{
		"sub": func(a, b int) int {
			return a - b
		},
	}).ParseFiles("./templates/index.html")
	if err != nil {
		errorPage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, w)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		errorPage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, w)
		return
	}
}
