package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	userRepository := UserRepository{db}
	return userRepository
}

func (uRepo UserRepository) Save(user User) error {
	now := time.Now()
	fmt.Println(user.Name.Value)
	userDTO := UserDTO{Name: user.Name.Value, MailAddress: user.MailAddress.Value, CreatedTime: now, UpdatedTime: now}
	result := uRepo.db.Create(&userDTO)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (uRepo UserRepository) FindByUserID(userID UserID) *User {
	var userDTO UserDTO
	uRepo.db.First(&userDTO, userID.Value)
	user := NewUser(userID, NewUserName(userDTO.Name), NewMailAddress(userDTO.MailAddress))
	return &user
}

func (UserRepository) FindByUserName(userName UserName) *User {
	return nil
}

func (UserRepository) FindByMailAddress(mailAddress MailAddress) *User {
	user := NewUser(NewUserID(1), NewUserName("takahiro"), mailAddress)
	return &user
}

func (UserRepository) Delete(user User) {
	return
}
