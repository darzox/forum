package service

import "forum/internal/model"

type Post interface {
	GetAllPosts() ([]model.PostRepresentation, error)
	CreatePost(heading string, text string, userId uint) (uint, error)
	GetPostById(postId uint) (*model.PostRepresentation, error)
}

type PostService struct {
	repo Post
}

func NewPostService(repo Post) *PostService {
	return &PostService{
		repo: repo,
	}
}

func (ps *PostService) GetAllPosts() ([]model.PostRepresentation, error) {
	return ps.repo.GetAllPosts()
}

func (ps *PostService) CreatePost(heading string, text string, userId uint) (uint, error) {
	return ps.repo.CreatePost(heading, text, userId)
}

func (ps *PostService) GetPostById(postId uint) (*model.PostRepresentation, error) {
	return ps.repo.GetPostById(postId)
}

