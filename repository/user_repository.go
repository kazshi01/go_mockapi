package repository

import (
	"errors"
	"mockapi/model"
)

type IUserRepository interface {
	GetUsers() model.AllUsers
	GetUserById(id int) (*model.User, error)
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
