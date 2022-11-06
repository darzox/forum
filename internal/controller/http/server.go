package controller

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"forum/internal/controller/http/handlers"
	"forum/internal/infrastructure/repository"
	"forum/internal/service"
	"forum/internal/storage"
)

func RunServer() {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", handlers.Index)
	// mux.HandleFunc("/err", err)
	mux.HandleFunc("/signin", handlers.SignIn)
	mux.HandleFunc("/logout", handlers.Logout)
	// mux.HandleFunc("/signup", handlers.SignUp)

	file, err := os.Create("database.db")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// repos := repository.NewUserRepository(db)
	repos1 := repository.NewRepository(db)

	// service := service.NewUserService(repos)
	service1 := service.NewService(repos1)

	control1 := handlers.NewContoller(service1)

	// control := handlers.NewUserController(&service1.UserServ)

	err = storage.CreateTables(db)
	if err != nil {
		log.Fatal(err)
	}
	mux.Handle("/signup", &control1.SignUp)

	mux.HandleFunc("/create-post", handlers.CreatePost)
	mux.HandleFunc("/post", handlers.Post)
	mux.HandleFunc("/create-comment", handlers.CreateComment)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
