package service

type Comment interface {
	CreateComment(userId, postId uint, text string) (uint, error)
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
