package handlers

import (
	"net/http"

	"forum/internal/service"
)

type Service interface {
	Registration
	Authorization
	Auth
	Leaving
	service.Post
}

type Controller struct {
	Index
	Logout
	SignUp
	SingIn
	Middleware
	CreatePost
	Post
}

func NewContoller(serv Service) *Controller {
	return &Controller{
		Index{},
		*CreateLogoutHandler(serv),
		*CreateSignUpHandler(serv),
		*CreateSignInHandler(serv),
		*CreateMiddleware(serv),
		*CreateCreatePostHandler(serv),
		*CreatePostHandler(serv),
	}
}

func (c *Controller) Run() error {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.Handle("/", c.AuthMiddleware(c.Index))
	// mux.HandleFunc("/err", err)
	// mux.HandleFunc("/signin", handlers.SignIn)
	mux.Handle("/logout", &c.Logout)
	// mux.HandleFunc("/signup", handlers.SignUp)
	mux.Handle("/signup", &c.SignUp)
	mux.Handle("/signin", &c.SingIn)
	mux.Handle("/create-post", c.AuthMiddleware(c.CreatePost))
	mux.Handle("/post", c.AuthMiddleware(&c.Post))
	mux.HandleFunc("/create-comment", CreateComment)

	server := http.Server{
		Addr:    "localhost:8081",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
