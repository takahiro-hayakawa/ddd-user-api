package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func main() {
	r := gin.Default()
	iUserRepository := NewUserRepository()
	userService := NewUserService(iUserRepository)
	userApplicationService := NewUserApplicationService(iUserRepository, userService)

	r.GET("/user/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, _ := strconv.Atoi(idStr)
		userDTO := userApplicationService.Find(id)

		c.JSON(200, gin.H{
			"user_id":      id,
			"name":         userDTO.Name,
			"maid_address": userDTO.MailAddress,
		})
	})

	r.POST("/user/", func(c *gin.Context) {
		userName := c.Param("name")
		userApplicationService.Register(userName)

		c.JSON(200, gin.H{
			"message": "success",
		})
	})

	r.Run(":3000")
}
