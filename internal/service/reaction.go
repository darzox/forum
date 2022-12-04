package service

type React interface {
	React(postOrComment string, userId uint, postId uint, positive bool) (uint, error)
}

type ReactService struct {
	repo React
}

func NewReacttService(repo React) *ReactService {
	return &ReactService{
		repo: repo,
	}
}

func (rs *ReactService) React(postOrComment string, userId uint, postId uint, positive bool) (uint, error) {
	return rs.repo.React(postOrComment, userId, postId, positive)
}
