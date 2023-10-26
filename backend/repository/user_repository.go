package repository

import "github.com/Akito-Fujihara/echo-go-react-js-todo/backend/model"

type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}
