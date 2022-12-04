package service

import "forum/internal/model"

type Comment interface {
	CreateComment(userId, postId uint, text string) (uint, error)
	GetAllCommentsByPostId(postId uint) ([]model.CommentRepresentation, error)
}

type CommentService struct {
	repo Comment
}

func NewCommentService(repo Comment) *CommentService {
	return &CommentService{
		repo: repo,
	}
}

func (cs *CommentService) CreateComment(userId, postId uint, text string) (uint, error) {
	return cs.repo.CreateComment(userId, postId, text)
}

func (cs *CommentService) GetAllCommentsByPostId(postId uint) ([]model.CommentRepresentation, error) {
	return cs.repo.GetAllCommentsByPostId(postId)
}
