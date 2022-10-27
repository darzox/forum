package main

import (
	"fmt"

	"forum/internal/app"
)

func main() {
	app.Run()
	fmt.Println("server started at http://localhost:8080/")
}
