package handlers

import (
	"net/http"
)

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

func (c *Controller) Run() error {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", Index)
	// mux.HandleFunc("/err", err)
	// mux.HandleFunc("/signin", handlers.SignIn)
	mux.HandleFunc("/logout", Logout)
	// mux.HandleFunc("/signup", handlers.SignUp)
	mux.Handle("/signup", &c.SignUp)
	mux.Handle("/signin", &c.SingIn)
	mux.HandleFunc("/create-post", CreatePost)
	mux.HandleFunc("/post", Post)
	mux.HandleFunc("/create-comment", CreateComment)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
