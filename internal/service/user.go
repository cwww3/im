package service

import (
	"errors"
	"github.com/cwww3/im/internal/model"
	"github.com/cwww3/im/internal/repository"
)

type UserService struct {
	*repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (us *UserService) GetUserList() ([]model.User, error) {
	return nil, nil
}

func (us *UserService) GetUserListExcept(user model.User) ([]model.User, error) {
	users, err := us.UserRepository.FindAll()
	if err != nil {
		return users, err
	}
	for i, u := range users {
		if u.Uid == user.Uid {
			users = append(users[:i], users[i+1:]...)
			break
		}
	}
	return users, err
}

func (us *UserService) GetContacts(user model.User) ([]model.User, error) {
	users, err := us.UserRepository.FindMessageContactsByOwnerUidOrderByMidDesc(user.Uid)
	if err != nil {
		return users, err
	}
	if len(users) == 0 {
		return users, err
	}
	return nil, err
}

func (us *UserService) Login(email, password string) (model.User, error) {
	user, err := us.UserRepository.FindByEmail(email)
	if err != nil {
		return model.User{}, errors.New("not found")
	}
	if user.Password != password {
		return model.User{}, errors.New("password error")
	}
	return user, nil
}
