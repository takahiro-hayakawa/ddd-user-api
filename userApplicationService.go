package main

import "fmt"

type UserApplicationService struct {
	userRepository IUserRepository
	userService    UserService
}

func NewUserApplicationService(userRepository IUserRepository, userService UserService) UserApplicationService {
	userApplicationService := UserApplicationService{userRepository, userService}
	return userApplicationService
}

func (userApplicationService UserApplicationService) Register(name string) error {
	userName := NewUserName(name)
	user := NewUserByUserName(userName)
	duplicatedUser := userApplicationService.userService.Exists(user)
	if duplicatedUser {
		return fmt.Errorf("user already exist")
	}
	fmt.Println(user.Name.Value)
	err := userApplicationService.userRepository.Save(user)
	if err != nil {
		return err
	}
	return nil
}

func (userApplicationService UserApplicationService) Find(userID int) User {
	targetUserID := NewUserID(userID)
	user := userApplicationService.userRepository.FindByUserID(targetUserID)
	return *user
}

func (userApplicationService UserApplicationService) Update(command UserUpdateCommand) (err error) {
	targetID := NewUserID(command.ID)
	user := userApplicationService.userRepository.FindByUserID(targetID)
	if user == nil {
		return fmt.Errorf("user not exist")
	}

	name := command.Name
	if name != "" {
		newUserName := NewUserName(name)
		user.ChangeName(newUserName)
		duplicatedUser := userApplicationService.userService.Exists(*user)
		if duplicatedUser {
			return fmt.Errorf("user already exist")
		}
	}

	mailAddress := command.MailAddress
	if mailAddress != "" {
		newMailAddress := NewMailAddress(mailAddress)
		user.ChangeMailAddress(newMailAddress)
	}
	userApplicationService.userRepository.Save(*user)
	return nil
}
