package handlers

import (
	"html/template"
	"net/http"

	"forum/internal/model"
)

type Registration interface {
	RegisterUser(user *model.User) error
}

type SignUp struct {
	service Registration
}

func CreateSignUpHandler(service Registration) *SignUp {
	return &SignUp{
		service: service,
	}
}

func (su *SignUp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		t, err := template.ParseFiles("./templates/signupPage.html")
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
	} else if r.Method == http.MethodPost {
		r.ParseForm()
		userInfo := r.PostForm
		user := model.User{
			Email:    userInfo["email"][0],
			Username: userInfo["username"][0],
			Password: userInfo["password"][0],
		}
		err := su.service.RegisterUser(&user)
		if err != nil {
			errorPage(err.Error(), http.StatusUnauthorized, w)
			return
		}
		http.Redirect(w, r, "/signin", http.StatusSeeOther)
		return
	}
	errorPage(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed, w)
}
