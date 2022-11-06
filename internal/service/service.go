package service

type Repository interface {
	UserRepository
}

type Service struct {
	userService
}

func NewService(repo Repository) *Service {
	return &Service{
		*NewUserService(repo),
	}
}
