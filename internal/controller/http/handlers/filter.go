package handlers

import (
	"fmt"
	"net/http"
	"strings"
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
	var err error
	user, ok := r.Context().Value("authorizedUser").(*model.User)
	r.ParseForm()
	filterBy := r.FormValue("filter_by")
	var filteredPosts []model.PostRepresentation
	if filterBy == "i_liked" || filterBy == "i_created" {
		filteredPosts, err = i.serv.PersonalFilter(filterBy, user.ID)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		filteredPosts, err = i.serv.FilterAllPosts((filterBy))
		if err != nil {
			fmt.Println(err)
		}
	}

	info := struct {
		User          *model.User
		Posts         []model.PostRepresentation
		HeadingFilter string
	}{
		User:          user,
		Posts:         filteredPosts,
		HeadingFilter: strings.Title(strings.Replace(filterBy, "_", " ", -1)) + " Posts",
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
