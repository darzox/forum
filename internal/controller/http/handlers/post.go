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
		Auth bool
	}{
		User: user,
		Post: post,
		Auth: true,
	}
	if !ok {
		fmt.Println("aaaaaaaaaa")
		info1 := struct {
			Post *model.PostRepresentation
			Auth bool
		}{
			Post: post,
			Auth: false,
		}
		t, err := template.New("post.html").Funcs(template.FuncMap{
			"sub": func(a, b int) int {
				return a - b
			},
		}).ParseFiles("./templates/post.html")
		if err != nil {
			fmt.Println(err)
		}
		err = t.Execute(w, info1)
		if err != nil {
			fmt.Println(err)
		}
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
