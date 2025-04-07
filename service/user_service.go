package service

import (
	"mockapi/model"
	"mockapi/repository"
)

type IUseServeice interface {
	GetUsers() model.AllUsers
	GetUserById(id int) (*model.User, error)
	Register(username string, password int) error
}

type UserService struct {
	ur repository.IUserRepository
}

func NewUserService(ur repository.IUserRepository) IUseServeice {
	return &UserService{ur}
}

func (us *UserService) GetUsers() model.AllUsers {
	users := us.ur.GetUsers()
	return users
}

func (us *UserService) GetUserById(id int) (*model.User, error) {
	user, err := us.ur.GetUserById(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *UserService) Register(username string, password int) error {
	err := us.ur.Register(username, password)

	return err
}
