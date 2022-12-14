package service

type Repository interface {
	RegisterUser
	LoginUser
	SessionCreator
	SessionChecker
	Post
	Comment
	React
}

type Service struct {
	RegisterUserService
	LoginUserService
	SessionCreateService
	SessionCheckService
	PostService
	CommentService
	ReactService
}

func NewService(repo Repository) *Service {
	return &Service{
		*NewRegisterUserService(repo),
		*NewLoginUserService(repo),
		*NewSessionCreateService(repo),
		*NewSessionCheckService(repo),
		*NewPostService(repo),
		*NewCommentService(repo),
		*NewReacttService(repo),
	}
}
