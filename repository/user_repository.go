package repository

import (
	"mockapi/model"
)

type IUserRepository interface {
	Say() string
}

type UserRepository struct {
	model *model.User
}

func NewUserRepository(model *model.User) IUserRepository {
	return &UserRepository{model}
}

func (ur *UserRepository) Say() string {
	return "Hello! My name is " + ur.model.Name
}
