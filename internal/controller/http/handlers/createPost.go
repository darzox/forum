package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"forum/internal/model"
	"forum/internal/service"
)

type CreatePost struct {
	serv service.Post
}

func CreateCreatePostHandler(serv service.Post) *CreatePost {
	return &CreatePost{
		serv: serv,
	}
}

func (cp CreatePost) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("authorizedUser").(*model.User)
	if r.Method == http.MethodGet {
		if !ok {
			errorPage(http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized, w)
			return
		}
		t, err := template.ParseFiles("./templates/createPostAuth.html")
		if err != nil {
			errorPage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, w)
			return
		}
		t.Execute(w, user.Username)
	}
	if r.Method == http.MethodPost {
		r.ParseForm()
		postInfo := r.PostForm
		post := model.PostRepresentation{
			Heading: postInfo["heading"][0],
			Text:    postInfo["text"][0],
		}
		postId, err := cp.serv.CreatePost(post.Heading, post.Text, user.ID)
		if err != nil {
			errorPage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, w)
			return
		}
		categories := postInfo["category"]
		for _, categoryId := range categories {
			categoryIdUint, _ := strconv.ParseUint(categoryId, 10, 32)
			_, err := cp.serv.AddCategoryToPost(uint(categoryIdUint), postId)
			if err != nil {
				errorPage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, w)
				return
			}
		}
		postIdString := strconv.FormatUint(uint64(postId), 10)
		http.Redirect(w, r, "/post?id="+postIdString, http.StatusSeeOther)
	}
}
