package main

type UserService struct {
	UserRepository IUserRepository
}

func NewUserService(userRepository IUserRepository) UserService {
	userService := UserService{userRepository}
	return userService
}

func (userService UserService) Exists(user User) bool {
	findUser := userService.UserRepository.FindByUserName(user.Name)
	if findUser != nil {
		return true
	}

	return false
}
