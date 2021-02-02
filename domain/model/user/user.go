package model

type User struct {
	ID          UserID
	Name        UserName
	MailAddress MailAddress
}

func NewUser(id UserID, name UserName, mailAddress MailAddress) User {
	user := User{id, name, mailAddress}
	return user
}

func NewUserByUserName(name UserName) User {
	ID := NewUserID(1)
	user := User{ID, name, NewMailAddress("")}
	return user
}

func (user *User) SetName(name UserName) {
	user.Name = name
}

func (user *User) GetName() UserName {
	return user.Name
}

func (user *User) SetMailAddress(mailAddress MailAddress) {
	user.MailAddress = mailAddress
}

func (user *User) GetMailAddress() MailAddress {
	return user.MailAddress
}

func (user *User) ChangeName(newName UserName) {
	user.SetName(newName)
}

func (user *User) ChangeMailAddress(newMailAddress MailAddress) {
	user.SetMailAddress(newMailAddress)
}
