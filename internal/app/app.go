package app

import (
	"fmt"

	"forum/internal/controller/http/handlers"
	"forum/internal/infrastructure/repository"
	"forum/internal/service"
)

func Run() error {
	db, err := repository.RunDb()
	if err != nil {
		fmt.Println("a")
		return err
	}
	defer db.Close()

	repos := repository.NewRepository(db)

	service := service.NewService(repos)

	control := handlers.NewContoller(service)

	if err := control.Run(); err != nil {
		return err
	}
	return nil
}
