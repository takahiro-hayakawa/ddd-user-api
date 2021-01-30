package main

type UserName struct {
	Value string
}

func NewUserName(value string) UserName {
	userName := UserName{value}
	return userName
}
