package main

type UserDTO struct {
	ID          int
	Name        string
	MailAddress string
}

func NewUserDTO(user User) UserDTO {
	userDTO := UserDTO{user.ID.Value, user.Name.Value, user.MailAddress.Value}
	return userDTO
}
