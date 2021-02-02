package model

type UserID struct {
	Value int
}

func NewUserID(value int) UserID {
	userID := UserID{value}
	return userID
}
