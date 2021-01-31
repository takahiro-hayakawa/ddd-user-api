package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"strconv"
)

func main() {
	db := gormConnect()
	defer db.Close()

	r := gin.Default()
	iUserRepository := NewUserRepository(db)
	userService := NewUserService(iUserRepository)
	userApplicationService := NewUserApplicationService(iUserRepository, userService)

	r.GET("/user/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, _ := strconv.Atoi(idStr)
		user := userApplicationService.Find(id)

		c.JSON(200, gin.H{
			"user_id":      id,
			"name":         user.Name.Value,
			"maid_address": user.MailAddress.Value,
		})
	})

	r.POST("/user/", func(c *gin.Context) {
		userName := c.Param("name")
		userApplicationService.Register(userName)

		c.JSON(200, gin.H{
			"message": "success",
		})
	})

	r.Run(":8000")
}

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "password"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "user_db"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}
