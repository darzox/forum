package service

type Repository interface {
	RegisterUser
	LoginUser
	SessionCreator
}

type Service struct {
	RegisterUserService
	LoginUserService
	SessionCreateService
}

func NewService(repo Repository) *Service {
	return &Service{
		*NewRegisterUserService(repo),
		*NewLoginUserService(repo),
		*NewSessionCreateService(repo),
	}
}
