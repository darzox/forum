package handlers

import (
	"html/template"
	"net/http"
	"time"

	"forum/internal/model"
)

type Authorization interface {
	LoginUser(user *model.User) (uint, bool, error)
	SessionCreate(user *model.User) (cookie string, err error)
}

type SingIn struct {
	service Authorization
}

func CreateSignInHandler(service Authorization) *SingIn {
	return &SingIn{
		service: service,
	}
}

func (si *SingIn) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, ok := r.Context().Value("authorizedUser").(*model.User)
	if ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if r.Method == http.MethodGet {
		t, err := template.ParseFiles("./templates/signinPage.html")
		if err != nil {
			errorPage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, w)
			return
		}
		err = t.Execute(w, nil)
		if err != nil {
			errorPage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, w)
			return
		}
		return
	}
	if r.Method == http.MethodPost {
		r.ParseForm()
		userInfo := r.PostForm
		for key := range r.PostForm {
			if !contains([]string{"username", "password"}, key) {
				errorPage(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, w)
				return
			}
		}
		user := model.User{
			Username: userInfo["username"][0],
			Password: userInfo["password"][0],
		}
		userId, userLogined, err := si.service.LoginUser(&user)
		if err != nil {
			errorPage(err.Error(), http.StatusUnauthorized, w)
			return
		}
		user.ID = userId
		if userLogined {
			// Create session
			cookie, err := si.service.SessionCreate(&user)
			if err != nil {
				errorPage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, w)
				return
			}
			cookieExpiresAt := time.Now().Add(600 * time.Second)
			http.SetCookie(w, &http.Cookie{
				Name:    "Session-token",
				Value:   cookie,
				Expires: cookieExpiresAt,
			})
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/signin", http.StatusSeeOther)
		}
		return
	}
	errorPage(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed, w)
}
