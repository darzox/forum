package service

import "forum/internal/model"

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

func (lus *LoginUserService) LoginUser(user *model.User) (bool, error) {
	userFromDB, err := lus.repo.GetUserByUsernameAndPassword(user)
	if err != nil {
		return false, err
	}

	if user.Username == userFromDB.Username && user.Password == userFromDB.Password {
		return true, nil
	}
	return false, nil
}
