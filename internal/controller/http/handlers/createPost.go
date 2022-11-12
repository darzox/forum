package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

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
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		t, err := template.ParseFiles("./templates/createPostAuth.html")
		if err != nil {
			fmt.Println(err)
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
		postIdString := strconv.FormatUint(uint64(postId), 10)
		if err != nil {
			fmt.Println(err)
		}
		http.Redirect(w, r, "/post?id="+postIdString, http.StatusSeeOther)
	}
}

// func CreatePost(w http.ResponseWriter, r *http.Request) {
// 	t, err := template.ParseFiles("./templates/createPost.html")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	t.Execute(w, nil)
// }
