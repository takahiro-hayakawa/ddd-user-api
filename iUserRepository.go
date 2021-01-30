package main

type IUserRepository interface {
	Save(user User)
	FindByUserID(userID UserID) *User
	FindByUserName(userName UserName) *User
	FindByMailAddress(mailAddress MailAddress) *User
	Delete(user User)
}
