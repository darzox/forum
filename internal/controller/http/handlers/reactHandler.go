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
	user, ok := r.Context().Value("authorizedUser").(*model.User)
	if !ok {
		errorPage(http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized, w)
		return
	}
	if r.Method != http.MethodPost {
		fmt.Println("a")
		errorPage(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed, w)
		return
	}
	r.ParseForm()
	for key := range r.PostForm {
		if !contains([]string{"positive", "postId", "reactTo", "commentId"}, key) {

			errorPage(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, w)
			return
		}
	}
	postOrComment := r.PostForm["reactTo"][0]
	if postOrComment == "post" {
		postIdString := r.Form["postId"][0]
		postId, err := strconv.Atoi(r.Form["postId"][0])
		if err != nil {
			errorPage(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, w)
			return
		}
		postIdUint := uint(postId)
		positive, err := strconv.ParseBool(r.Form["positive"][0])
		if err != nil {
			errorPage(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, w)
			return
		}
		_, err = re.serv.React(postOrComment, user.ID, postIdUint, positive)
		if err != nil {
			errorPage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, w)
			return
		}
		http.Redirect(w, r, "/post?id="+postIdString, http.StatusSeeOther)
		return
	}
	if postOrComment == "comment" {
		postIdString := r.Form["postId"][0]
		commentId, err := strconv.Atoi(r.Form["commentId"][0])
		if err != nil {
			errorPage(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, w)
			return
		}
		commentIdUint := uint(commentId)
		positive, err := strconv.ParseBool(r.Form["positive"][0])
		if err != nil {
			fmt.Println("a")
			errorPage(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, w)
			return
		}
		_, err = re.serv.React(postOrComment, user.ID, commentIdUint, positive)
		if err != nil {
			errorPage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, w)
			return
		}
		http.Redirect(w, r, "/post?id="+postIdString, http.StatusSeeOther)
		return
	}
}
