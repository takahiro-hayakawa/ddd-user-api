package main

type UserRepository struct {
}

func NewUserRepository() UserRepository {
	userRepository := UserRepository{}
	return userRepository
}

func (UserRepository) Save(user User) {
	return
}

func (UserRepository) FindByUserID(userID UserID) *User {
	user := NewUser(userID, NewUserName("takahiro"), NewMailAddress("hoge@aaa.com"))
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
