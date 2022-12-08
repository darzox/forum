package main

import (
	"fmt"
	"log"

	"forum/internal/app"
)

func main() {
	fmt.Println("server started at http://localhost:8081/")
	err := app.Run()
	if err != nil {
		log.Fatalf(err.Error(), err)
	}
}
