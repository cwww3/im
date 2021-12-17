package service

import "github.com/cwww3/im/internal/model"

type IUserService interface {
	GetUserList() ([]model.User,error)
	GetUserListExcept(model.User) ([]model.User,error)
	GetContacts(model.User) ([]model.User,error)
	Login(email,password string) (model.User,error)
}

