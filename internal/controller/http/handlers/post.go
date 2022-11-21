package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"forum/internal/model"
	"forum/internal/service"
)

type Post struct {
	serv service.Post
}

func CreatePostHandler(serv Service) *Post {
	return &Post{
		serv: serv,
	}
}

func (p *Post) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("authorizedUser").(*model.User)
	fmt.Println(user)
	postIdSting := r.URL.Query().Get("id")
	postId64, _ := strconv.ParseUint(postIdSting, 10, 32)
	postId := uint(postId64)

	post, err := p.serv.GetPostById(postId)
	if err != nil {
		fmt.Println(err)
	}
	info := struct {
		User *model.User
		Post *model.PostRepresentation
	}{
		User: user,
		Post: post,
	}
	if !ok {
		t, err := template.New("post.html").Funcs(template.FuncMap{
			"sub": func(a, b int) int {
				return a - b
			},
		}).ParseFiles("./templates/post.html")
		if err != nil {
			fmt.Println()
		}
		t.Execute(w, info)
		return
	}
	fmt.Println(info.User, info.Post)
	// function inside template
	t, err := template.New("post.html").Funcs(template.FuncMap{
		"sub": func(a, b int) int {
			return a - b
		},
	}).ParseFiles("./templates/post.html")
	if err != nil {
		fmt.Println()
	}
	t.Execute(w, info)
}
