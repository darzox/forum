package handlers

import (
	"html/template"
	"net/http"
	"strings"

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

func (f Filter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorPage(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed, w)
		return
	}
	var err error
	user, ok := r.Context().Value("authorizedUser").(*model.User)
	if !ok {
		r.ParseForm()
		if _, ok := r.Form["filter_by"]; !ok {
			errorPage(http.StatusText(http.StatusNotFound), http.StatusNotFound, w)
			return
		}
		value := r.Form["filter_by"]
		if !contains([]string{"oldest", "recent", "most_disliked", "most_liked", "discussions", "questions", "ideas", "articles", "events", "issues", "i_liked", "i_created"}, value[0]) {
			errorPage(http.StatusText(http.StatusNotFound), http.StatusNotFound, w)
			return
		}
		filterBy := r.FormValue("filter_by")
		var filteredPosts []model.PostRepresentation
		filteredPosts, err = f.serv.FilterAllPosts((filterBy))
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
			Posts:         filteredPosts,
			HeadingFilter: strings.Title(strings.Replace(filterBy, "_", " ", -1)) + " Posts",
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
		err = t.Execute(w, info)
		if err != nil {
			errorPage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, w)
			return
		}
	}
	r.ParseForm()
	if _, ok := r.Form["filter_by"]; !ok {
		errorPage(http.StatusText(http.StatusNotFound), http.StatusNotFound, w)
		return
	}
	value := r.Form["filter_by"]
	if !contains([]string{"oldest", "recent", "most_disliked", "most_liked", "discussions", "questions", "ideas", "articles", "events", "issues", "i_liked", "i_created"}, value[0]) {
		errorPage(http.StatusText(http.StatusNotFound), http.StatusNotFound, w)
		return
	}
	filterBy := r.FormValue("filter_by")
	var filteredPosts []model.PostRepresentation
	if filterBy == "i_liked" || filterBy == "i_created" {
		filteredPosts, err = f.serv.PersonalFilter(filterBy, user.ID)
		if err != nil {
			errorPage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, w)
			return
		}
	} else {
		filteredPosts, err = f.serv.FilterAllPosts((filterBy))
		if err != nil {
			errorPage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, w)
			return
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
	err = t.Execute(w, info)
	if err != nil {
		errorPage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, w)
		return
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
