package handlers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"forum/internal/model"
)

type Authorization interface {
	LoginUser(user *model.User) (bool, error)
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
	if r.Method == http.MethodGet {
		t, err := template.ParseFiles("./templates/signinPage.html")
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
			Username: userInfo["username"][0],
			Password: userInfo["password"][0],
		}
		userLogined, err := si.service.LoginUser(&user)
		if err != nil {
			http.Redirect(w, r, "/err", http.StatusSeeOther)
		}
		if userLogined {
			// Create session
			http.Redirect(w, r, "err", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/signin", http.StatusSeeOther)
		}
	}
}

// func SignIn(w http.ResponseWriter, r *http.Request) {
// 	t, err := template.ParseFiles("./templates/signinPage.html")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	t.Execute(w, nil)
// }
