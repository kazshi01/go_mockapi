package service

import (
	"mockapi/repository"
)

type IUseServeice interface {
	SayPlus() string
}

type UserService struct {
	ur repository.IUserRepository
}

func NewUserService(ur repository.IUserRepository) IUseServeice {
	return &UserService{ur}
}

func (us *UserService) SayPlus() string {
	message := us.ur.Say()
	return message + "\nPlus extra functionality!"
}
