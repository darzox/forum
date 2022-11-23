package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"forum/internal/model"
	"forum/internal/service"
)

type CreateComment struct {
	serv service.Comment
}

func CreateCommentHandler(serv Service) *CreateComment {
	return &CreateComment{
		serv: serv,
	}
}

func (cc CreateComment) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.Redirect(w, r, "/err", http.StatusMethodNotAllowed)
	}
	r.ParseForm()
	commentInfo := r.PostForm
	postIdString := commentInfo["postId"][0]
	postId64, _ := strconv.ParseUint(postIdString, 10, 32)
	postId := uint(postId64)
	userIdSting := commentInfo["userId"][0]
	userId64, _ := strconv.ParseUint(userIdSting, 10, 32)
	userId := uint(userId64)
	comment := model.CommentRepresentation{
		PostId: postId,
		UserId: userId,
		Text:   commentInfo["comment"][0],
	}
	_, err := cc.serv.CreateComment(comment.UserId, comment.PostId, comment.Text)
	if err != nil {
		fmt.Println(err)
	}
	http.Redirect(w, r, "/post?id="+postIdString, http.StatusSeeOther)
}
