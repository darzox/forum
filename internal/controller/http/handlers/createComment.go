package handlers

import (
	"log"
	"net/http"
	"text/template"
)

func CreateComment(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/signinPage.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, nil)
}
