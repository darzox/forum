package handlers

type Service interface {
	Registration
}

type Controller struct {
	SignUp
}

func NewContoller(serv Service) *Controller {
	return &Controller{
		*CreateSignUpHandler(serv),
	}
}
