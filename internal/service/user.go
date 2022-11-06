package service

import (
	"fmt"

	"forum/internal/model"
)

type UserRepository interface {
	CreateUser(user *model.User) error
	GetUser(user *model.User) (*model.User, error)
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *userService {
	return &userService{
		repo: repo,
	}
}

func (us *userService) CreateUser(user *model.User) error {
	userFromDB, err := us.repo.GetUser(user)
	if err != nil {
		return err
	}
	if userFromDB.Email == user.Email || userFromDB.Username == user.Username {
		return fmt.Errorf("user already exists")
	}
	return us.repo.CreateUser(user)
}
