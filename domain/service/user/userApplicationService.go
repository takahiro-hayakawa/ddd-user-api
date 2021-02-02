package service

import (
	"fmt"
	mUser "github.com/takahiro-hayakawa/user-api-server/domain/model/user"
)

type UserApplicationService struct {
	userRepository mUser.IUserRepository
	userService    UserService
}

func NewUserApplicationService(userRepository mUser.IUserRepository, userService UserService) UserApplicationService {
	userApplicationService := UserApplicationService{userRepository, userService}
	return userApplicationService
}

func (userApplicationService UserApplicationService) Register(name string) error {
	userName := mUser.NewUserName(name)
	user := mUser.NewUserByUserName(userName)
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

func (userApplicationService UserApplicationService) Find(userID int) mUser.User {
	targetUserID := mUser.NewUserID(userID)
	user := userApplicationService.userRepository.FindByUserID(targetUserID)
	return *user
}

func (userApplicationService UserApplicationService) Update(command UserUpdateCommand) error {
	targetID := mUser.NewUserID(command.ID)
	user := userApplicationService.userRepository.FindByUserID(targetID)
	if user == nil {
		return fmt.Errorf("user not exist")
	}

	name := command.Name
	if name != "" {
		newUserName := mUser.NewUserName(name)
		user.ChangeName(newUserName)
		duplicatedUser := userApplicationService.userService.Exists(*user)
		if duplicatedUser {
			return fmt.Errorf("user already exist")
		}
	}

	mailAddress := command.MailAddress
	if mailAddress != "" {
		newMailAddress := mUser.NewMailAddress(mailAddress)
		user.ChangeMailAddress(newMailAddress)
	}
	err := userApplicationService.userRepository.Save(*user)
	if err != nil {
		return err
	}
	return nil
}
