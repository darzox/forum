package handlers

import (
	"fmt"
	"html/template"
	"log"
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
	cookieFromClient := r.Header.Get("Cookie")
	cookieFromClient = strings.ReplaceAll(cookieFromClient, "Session-token=", "")
	user, err := l.service.UserBySession(cookieFromClient)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/err", http.StatusSeeOther)
	}
	err = l.service.DeleteSession(user.ID)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/err", http.StatusSeeOther)
	}
	t, err := template.New("index.html").Funcs(template.FuncMap{
		"sub": func(a, b int) int {
			return a - b
		},
	}).ParseFiles("./templates/index.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, nil)
}
