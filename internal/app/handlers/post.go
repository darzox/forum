package handlers

import (
	"log"
	"net/http"
	"text/template"
)

func Post(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/post.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, nil)
}
