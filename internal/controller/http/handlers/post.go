package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"forum/internal/model"
)

type Post struct {
	serv Service
}

func CreatePostHandler(serv Service) *Post {
	return &Post{
		serv: serv,
	}
}

func (p *Post) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorPage(http.StatusText(http.StatusNotFound), http.StatusNotFound, w)
		return
	}
	user, ok := r.Context().Value("authorizedUser").(*model.User)
	postIdSting := r.URL.Query().Get("id")
	allPosts, err := p.serv.GetAllPosts()
	if err != nil {
		errorPage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, w)
		return
	}
	postId64, err := strconv.ParseUint(postIdSting, 10, 32)
	if err != nil || postId64 == 0 || int(postId64) > len(allPosts) {
		errorPage(http.StatusText(http.StatusNotFound), http.StatusNotFound, w)
		return
	}
	postId := uint(postId64)

	post, err := p.serv.GetPostById(postId)
	if err != nil {
		errorPage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, w)
		return
	}
	comments, err := p.serv.GetAllCommentsByPostId(postId)
	if err != nil {
		errorPage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, w)
		return
	}
	info := struct {
		User     *model.User
		Post     *model.PostRepresentation
		Auth     bool
		Comments []model.CommentRepresentation
	}{
		User:     user,
		Post:     post,
		Auth:     true,
		Comments: comments,
	}
	if !ok {
		info1 := struct {
			Post     *model.PostRepresentation
			Auth     bool
			Comments []model.CommentRepresentation
		}{
			Post:     post,
			Auth:     false,
			Comments: comments,
		}
		t, err := template.New("post.html").Funcs(template.FuncMap{
			"sub": func(a, b int) int {
				return a - b
			},
		}).ParseFiles("./templates/post.html")
		if err != nil {
			errorPage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, w)
			return
		}
		err = t.Execute(w, info1)
		if err != nil {
			errorPage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, w)
			return
		}
		return
	}
	// function inside template
	t, err := template.New("post.html").Funcs(template.FuncMap{
		"sub": func(a, b int) int {
			return a - b
		},
	}).ParseFiles("./templates/post.html")
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
