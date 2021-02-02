package service

import "time"

type UserDTO struct {
	ID          int `gorm:"primary_key;auto_increment;not_null"`
	Name        string
	MailAddress string
	CreatedTime time.Time
	UpdatedTime time.Time
}

func (UserDTO) TableName() string {
	return "user"
}
