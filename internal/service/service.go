package service

type Repository interface {
	RegisterUser
	LoginUser
	SessionCreator
	SessionChecker
}

type Service struct {
	RegisterUserService
	LoginUserService
	SessionCreateService
	SessionCheckService
}

func NewService(repo Repository) *Service {
	return &Service{
		*NewRegisterUserService(repo),
		*NewLoginUserService(repo),
		*NewSessionCreateService(repo),
		*NewSessionCheckService(repo),
	}
}
