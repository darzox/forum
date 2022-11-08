package handlers

type Service interface {
	Registration
	Authorization
}

type Controller struct {
	SignUp
	SingIn
}

func NewContoller(serv Service) *Controller {
	return &Controller{
		*CreateSignUpHandler(serv),
		*CreateSignInHandler(serv),
	}
}
