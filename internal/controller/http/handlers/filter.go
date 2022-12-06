package handlers

import (
	"fmt"
	"net/http"
	"text/template"

	"forum/internal/model"
	"forum/internal/service"
)

type FilterInterface interface {
	GetAllPosts() ([]model.PostRepresentation, error)
}

type Filter struct {
	serv service.Post
}

func CreateFilterHandler(serv service.Post) *Filter {
	return &Filter{
		serv: serv,
	}
}

func (i Filter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("authorizedUser").(*model.User)
	r.ParseForm()
	filterBy := r.FormValue("filter_by")
	filteredPosts, err := i.serv.FilterAllPosts((filterBy))
	info := struct {
		User  *model.User
		Posts []model.PostRepresentation
	}{
		User:  user,
		Posts: filteredPosts,
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