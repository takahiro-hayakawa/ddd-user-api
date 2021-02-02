package service

import (
	mUser "github.com/takahiro-hayakawa/user-api-server/domain/model/user"
)

type UserService struct {
	UserRepository mUser.IUserRepository
}

func NewUserService(userRepository mUser.IUserRepository) UserService {
	userService := UserService{userRepository}
	return userService
}

func (userService UserService) Exists(user mUser.User) bool {
	findUser := userService.UserRepository.FindByUserName(user.Name)
	if findUser != nil {
		return true
	}

	return false
}
