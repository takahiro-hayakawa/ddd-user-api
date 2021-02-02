package model

type MailAddress struct {
	Value string
}

func NewMailAddress(address string) MailAddress {
	mailAddress := MailAddress{address}
	return mailAddress
}
