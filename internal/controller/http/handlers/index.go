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
	info := struct {
		User  *model.User
		Posts []model.PostRepresentation
	}{
		User:  user,
		Posts: allposts,
	}
	if err != nil {
		fmt.Println(err)
	}
	if !ok {
		t, err := template.New("index.html").Funcs(template.FuncMap{
			"sub": func(a, b int) int {
				return a - b
			},
		}).ParseFiles("./templates/index.html")
		if err != nil {
			fmt.Println(err)
		}
		t.Execute(w, info)
		return
	}
	t, err := template.New("indexAuthorized.html").Funcs(template.FuncMap{
		"sub": func(a, b int) int {
			return a - b
		},
	}).ParseFiles("./templates/indexAuthorized.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, info)
}
