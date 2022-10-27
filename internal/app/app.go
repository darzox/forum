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
	// mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", handlers.SignUp)

	// mux.HandleFunc("/create-post", createPost)
	// mux.HandleFunc("/post", post)
	// mux.HandleFunc("/create-comment", createComment)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
