package handlers

import (
	"fmt"
	"net/http"
	"text/template"

	"forum/internal/model"
)

type Index struct{}

func (i Index) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("authorizedUser").(*model.User)
	if !ok {
		t, err := template.ParseFiles("./templates/index.html")
		if err != nil {
			fmt.Println(err)
		}
		t.Execute(w, nil)
		return
	}
	t, err := template.ParseFiles("./templates/indexAuthorized.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, user.Username)
}
