package handlers

import (
	"log"
	"net/http"
	"text/template"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/signupPage.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, nil)
}
