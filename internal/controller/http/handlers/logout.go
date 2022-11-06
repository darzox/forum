package handlers

import (
	"log"
	"net/http"
	"text/template"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/index")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, nil)
}
