package handlers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	"forum/internal/model"
)

type Authorization interface {
	LoginUser(user *model.User) (bool, error)
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
			fmt.Println(err)
			http.Redirect(w, r, "/err", http.StatusSeeOther)
		}
		if userLogined {
			// Create session
			cookie, err := si.service.SessionCreate(&user)
			if err != nil {
				fmt.Println(err)
				http.Redirect(w, r, "err", http.StatusSeeOther)
			}
			cookieExpiresAt := time.Now().Add(600 * time.Second)
			http.SetCookie(w, &http.Cookie{
				Name:    "session-token",
				Value:   cookie,
				Expires: cookieExpiresAt,
			})
			http.Redirect(w, r, "/", http.StatusSeeOther)
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
