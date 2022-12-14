package service

import (
	"fmt"

	"forum/internal/model"
)

type RegisterUser interface {
	CreateUser(user *model.User) error
	GetUser(user *model.User) (*model.User, error)
}

type RegisterUserService struct {
	repo RegisterUser
}

func NewRegisterUserService(repo RegisterUser) *RegisterUserService {
	return &RegisterUserService{
		repo: repo,
	}
}

func (rus *RegisterUserService) RegisterUser(user *model.User) error {
	userFromDB, err := rus.repo.GetUser(user)
	if err != nil {
		return fmt.Errorf("Internal Server Error")
	}
	if userFromDB.Email == user.Email || userFromDB.Username == user.Username {
		return fmt.Errorf("User with this email or username already exists")
	}
	return rus.repo.CreateUser(user)
}


