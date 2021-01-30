package main

type UserUpdateCommand struct {
	ID          int
	Name        string
	MailAddress string
}

func NewUserUpdateCommand(id int, name string, mailAddress string) UserUpdateCommand {
	userUpdateCommand := UserUpdateCommand{id, name, mailAddress}
	return userUpdateCommand
}
