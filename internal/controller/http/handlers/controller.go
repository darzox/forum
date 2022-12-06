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
	service.Comment
	service.React
}

type Controller struct {
	Index
	Logout
	SignUp
	SingIn
	Middleware
	CreatePost
	Post
	CreateComment
	React
	Filter
}

func NewContoller(serv Service) *Controller {
	return &Controller{
		*CreateIndexHandler(serv),
		*CreateLogoutHandler(serv),
		*CreateSignUpHandler(serv),
		*CreateSignInHandler(serv),
		*CreateMiddleware(serv),
		*CreateCreatePostHandler(serv),
		*CreatePostHandler(serv),
		*CreateCommentHandler(serv),
		*CreateReactHandler(serv),
		*CreateFilterHandler(serv),
	}
}

func (c *Controller) Run() error {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.Handle("/likeup", c.AuthMiddleware(c.React))
	mux.Handle("/likedown", c.AuthMiddleware(c.React))
	mux.Handle("/", c.AuthMiddleware(&c.Index))
	mux.Handle("/logout", &c.Logout)
	mux.Handle("/signup", &c.SignUp)
	mux.Handle("/signin", &c.SingIn)
	mux.Handle("/create-post", c.AuthMiddleware(c.CreatePost))
	mux.Handle("/post", c.AuthMiddleware(&c.Post))
	mux.Handle("/create-comment", &c.CreateComment)
	mux.Handle("/filter", c.AuthMiddleware(&c.Filter))

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
