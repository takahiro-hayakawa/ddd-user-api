package main

import "fmt"

func main() {
	iUserRepository := NewUserRepository()
	userService := NewUserService(iUserRepository)
	userApplicationService := NewUserApplicationService(iUserRepository, userService)
	err := userApplicationService.Register("takahiro")
	if err != nil {
		panic(err)
	}
	userDTO := userApplicationService.Find(1)
	fmt.Printf("userID:%d\n", userDTO.ID)
	fmt.Printf("userID:%s\n", userDTO.Name)
	fmt.Printf("userID:%s\n", userDTO.MailAddress)
}
