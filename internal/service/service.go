package service

type Repository interface {
	RegisterUser
	LoginUser
}

type Service struct {
	RegisterUserService
	LoginUserService
}

func NewService(repo Repository) *Service {
	return &Service{
		*NewRegisterUserService(repo),
		*NewLoginUserService(repo),
	}
}
