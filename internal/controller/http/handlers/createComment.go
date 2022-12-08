package handlers

import (
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
		errorPage(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed, w)
		return
	}
	r.ParseForm()
	commentInfo := r.PostForm
	postIdString := commentInfo["postId"][0]
	postId64, err := strconv.ParseUint(postIdString, 10, 32)
	if err != nil {
		errorPage(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, w)
		return
	}
	postId := uint(postId64)
	userIdSting := commentInfo["userId"][0]
	userId64, err := strconv.ParseUint(userIdSting, 10, 32)
	if err != nil {
		errorPage(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, w)
		return
	}
	userId := uint(userId64)
	comment := model.CommentRepresentation{
		PostId: postId,
		UserId: userId,
		Text:   commentInfo["comment"][0],
	}
	_, err = cc.serv.CreateComment(comment.UserId, comment.PostId, comment.Text)
	if err != nil {
		errorPage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, w)
		return
	}
	http.Redirect(w, r, "/post?id="+postIdString, http.StatusSeeOther)
}
