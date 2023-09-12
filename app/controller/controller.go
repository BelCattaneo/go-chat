package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	fmt.Println("Create User")
	username := c.PostForm("username")

	c.JSON(http.StatusOK, gin.H{
		"status":   "success",
		"username": username,
	})
}
