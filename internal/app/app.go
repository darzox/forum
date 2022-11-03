package app

import (
	"net/http"

	"forum/internal/app/handlers"
)

func Run() {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", handlers.Index)
	// mux.HandleFunc("/err", err)
	mux.HandleFunc("/login", handlers.SignIn)
	mux.HandleFunc("/logout", handlers.Logout)
	mux.HandleFunc("/signup", handlers.SignUp)

	mux.HandleFunc("/create-post", handlers.CreatePost)
	mux.HandleFunc("/post", handlers.Post)
	mux.HandleFunc("/create-comment", handlers.CreateComment)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
