package repository

import (
	"errors"
	"mockapi/model"
)

type IUserRepository interface {
	GetUsers() model.AllUsers
	GetUserById(id int) (*model.User, error)
	Register(username string, password int) error
}

type UserRepository struct{}

func NewUserRepository() IUserRepository {
	return &UserRepository{}
}

func (ur *UserRepository) GetUsers() model.AllUsers {
	return model.Users
}

func (ur *UserRepository) GetUserById(id int) (*model.User, error) {
	for _, user := range model.Users {
		if id == user.Id {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (ur *UserRepository) Register(username string, password int) error {
	_, exist := model.Registers[username]
	if exist {
		return errors.New("すでに登録済みです")
	}

	model.Registers[username] = model.Register{Username: username, Password: password}
	return nil
}
