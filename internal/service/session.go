package service

import (
	"forum/internal/model"

	"github.com/google/uuid"
)

type SessionCreator interface {
	CreateSession(cookie string, userId uint) error
}

type SessionCreateService struct {
	repo SessionCreator
}

func NewSessionCreateService(repo SessionCreator) *SessionCreateService {
	return &SessionCreateService{
		repo: repo,
	}
}

func (scs *SessionCreateService) SessionCreate(user *model.User) (cookie string, err error) { 
	cookieString := uuid.New().String()
	err = scs.repo.CreateSession(cookieString, user.ID)
	if err != nil {
		return "", err
	}
	return cookieString, nil
}
