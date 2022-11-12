package handlers

import (
	"fmt"
	"net/http"
	"text/template"

	"forum/internal/model"
)

type IndexInterface interface {
	GetAllPosts() ([]model.PostRepresentation, error)
}

type Index struct {
	service IndexInterface
}

func CreateIndexHandler(serv IndexInterface) *Index {
	return &Index{
		service: serv,
	}
}

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
	// allposts, err := i.service.GetAllPosts()
	// fmt.Println(allposts)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	t, err := template.ParseFiles("./templates/indexAuthorized.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, user.Username)
}
