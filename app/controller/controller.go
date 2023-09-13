package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/BelCattaneo/go-chat/app/database"
	"github.com/BelCattaneo/go-chat/app/model"
)

func CreateUser(c *gin.Context) {
	fmt.Println("Create User")
	fmt.Println("context is:")
	fmt.Println(c)
	fmt.Println(c.Params)
	username := c.PostForm("username")
	conn, _ := database.ConnectDB()
	db, _ := conn.DB()
	defer db.Close()
	conn.Create(&model.User{Username: username})
	c.JSON(http.StatusOK, gin.H{
		"status":   "success",
		"username": username,
	})
}
