package handlers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"forum/internal/model"
)

type Registration interface {
	CreateUser(user *model.User) error
}

type SignUp struct {
	service Registration
}

func CreateSignUpHandler(service Registration) *SignUp {
	return &SignUp{
		service: service,
	}
}

func (uc *SignUp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		t, err := template.ParseFiles("./templates/signupPage.html")
		if err != nil {
			log.Fatal(err)
		}
		t.Execute(w, nil)
	}
	if r.Method == http.MethodPost {
		r.ParseForm()
		userInfo := r.PostForm
		fmt.Println(userInfo)
		user := model.User{
			Email:    userInfo["email"][0],
			Username: userInfo["username"][0],
			Password: userInfo["password"][0],
		}
		err := uc.service.CreateUser(&user)
		if err != nil {
			http.Redirect(w, r, "/err", http.StatusSeeOther)
		}
		http.Redirect(w, r, "/signin", http.StatusSeeOther)
	}
}
