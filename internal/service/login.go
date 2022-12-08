package service

import (
	"fmt"

	"forum/internal/model"
)

type LoginUser interface {
	GetUserByUsernameAndPassword(user *model.User) (*model.User, error)
}

type LoginUserService struct {
	repo LoginUser
}

func NewLoginUserService(repo LoginUser) *LoginUserService {
	return &LoginUserService{
		repo: repo,
	}
}

func (lus *LoginUserService) LoginUser(user *model.User) (uint, bool, error) {
	userFromDB, err := lus.repo.GetUserByUsernameAndPassword(user)
	if err != nil {
		return 0, false, fmt.Errorf("Internal server error")
	}
	if user.Username == userFromDB.Username && user.Password == userFromDB.Password {
		return userFromDB.ID, true, nil
	}
	return 0, false, fmt.Errorf("Username or password is not correct")
}
