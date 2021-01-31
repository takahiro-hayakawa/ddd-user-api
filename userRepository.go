package main

import "github.com/jinzhu/gorm"

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	userRepository := UserRepository{db}
	return userRepository
}

func (UserRepository) Save(user User) {
	return
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
