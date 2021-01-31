package main

import "time"

type UserDTO struct {
	ID          int `gorm:"AUTO_INCREMENT"`
	Name        string
	MailAddress string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}

func (UserDTO) TableName() string {
	return "user"
}
