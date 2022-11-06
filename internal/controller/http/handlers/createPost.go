package handlers

import (
	"log"
	"net/http"
	"text/template"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/createPost.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, nil)
}

