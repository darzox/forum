package main

import (
	"fmt"

	"forum/internal/app"
	"forum/internal/storage"
)

func main() {
	storage.RunDb()
	app.Run()
	fmt.Println("server started at http://localhost:8080/")
}
