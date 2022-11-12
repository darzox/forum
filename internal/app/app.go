package app

import (
	"fmt"

	"forum/internal/controller/http/handlers"
	"forum/internal/infrastructure/repository"
	"forum/internal/service"
)

func Run() {
	db, err := repository.RunDb()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	repos1 := repository.NewRepository(db)

	service1 := service.NewService(repos1)

	control1 := handlers.NewContoller(service1)

	if err := control1.Run(); err != nil {
		fmt.Println(err)
	}
}
