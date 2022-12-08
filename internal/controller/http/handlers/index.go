package handlers

import (
	"html/template"
	"net/http"

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
	if r.URL.Path != "/" {
		errorPage(http.StatusText(http.StatusNotFound), http.StatusNotFound, w)
		return
	}
	user, ok := r.Context().Value("authorizedUser").(*model.User)
	allposts, err := i.serv.GetAllPosts()
	if err != nil {
		errorPage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, w)
		return
	}
	info := struct {
		User          *model.User
		Posts         []model.PostRepresentation
		HeadingFilter string
	}{
		User:          user,
		Posts:         allposts,
		HeadingFilter: "Latest Posts",
	}
	if !ok {
		t, err := template.New("index.html").Funcs(template.FuncMap{
			"sub": func(a, b int) int {
				return a - b
			},
		}).ParseFiles("./templates/index.html")
		if err != nil {
			errorPage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, w)
			return
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
		errorPage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, w)
		return
	}
	t.Execute(w, info)
}
