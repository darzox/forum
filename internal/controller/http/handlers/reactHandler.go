package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"forum/internal/model"
)

type React struct {
	serv Service
}

func CreateReactHandler(serv Service) *React {
	return &React{
		serv: serv,
	}
}

func (re React) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, _ := r.Context().Value("authorizedUser").(*model.User)
	if r.Method != http.MethodGet {
		postIdString := r.Form["postId"][0]
		http.Redirect(w, r, "/post?id="+postIdString, http.StatusSeeOther)
	}
	r.ParseForm()
	postOrComment := r.Form["reactTo"][0]
	if postOrComment == "post" {
		postIdString := r.Form["postId"][0]
		postId, err := strconv.Atoi(r.Form["postId"][0])
		postIdUint := uint(postId)
		if err != nil {
			fmt.Println(err)
		}
		positive, err := strconv.ParseBool(r.Form["positive"][0])
		if err != nil {
			fmt.Println(err)
		}
		re.serv.React(postOrComment, user.ID, postIdUint, positive)
		http.Redirect(w, r, "/post?id="+postIdString, http.StatusSeeOther)
	}
	if postOrComment == "comment" {
		postIdString := r.Form["postId"][0]
		postId, err := strconv.Atoi(r.Form["postId"][0])
		postIdUint := uint(postId)
		if err != nil {
			fmt.Println(err)
		}
		positive, err := strconv.ParseBool(r.Form["positive"][0])
		if err != nil {
			fmt.Println(err)
		}
		re.serv.React(postOrComment, user.ID, postIdUint, positive)
		http.Redirect(w, r, "/post?id="+postIdString, http.StatusSeeOther)
	}
}
