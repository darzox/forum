package handlers

import (
	"fmt"
	"net/http"
	"text/template"

	"forum/internal/model"
	"forum/internal/service"
)

type IndexInterface interface {
	GetAllPosts() ([]model.PostRepresentation, error)
}

type Index struct {
	serv service.Post
}

func CreateIndexHandler(serv service.Post) *Index {
	return &Index{
		serv: serv,
	}
}

func (i Index) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("authorizedUser").(*model.User)
	allposts, err := i.serv.GetAllPosts()
	fmt.Println(allposts)
	if err != nil {
		fmt.Println(err)
	}
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
