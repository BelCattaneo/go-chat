package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/BelCattaneo/go-chat/app/database"
	"github.com/BelCattaneo/go-chat/app/model"
)

func CreateUser(c *gin.Context) {
	fmt.Println("Create User")

	var input struct {
		Username string `json:"username"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		// Handle JSON parsing error
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	username := input.Username

	conn, _ := database.ConnectDB()
	db, _ := conn.DB()
	defer db.Close()
	conn.Create(&model.User{Username: username})

	c.JSON(http.StatusOK, gin.H{
		"status":   "success",
		"username": username,
	})
}

func CreateRoom(c *gin.Context) {
	fmt.Println("Create Room")

	var input struct {
		Roomname string `json:"roomname"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		// Handle JSON parsing error
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	roomname := input.Roomname

	conn, _ := database.ConnectDB()
	db, _ := conn.DB()
	defer db.Close()
	conn.Create(&model.Room{Roomname: roomname})

	c.JSON(http.StatusOK, gin.H{
		"status":   "success",
		"roomname": roomname,
	})
}

func EnterRoom(c *gin.Context) {
	fmt.Println("Enter Room")

	var input struct {
		RoomId string `json:"roomId"`
		UserId string `json:"userId"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		// Handle JSON parsing error
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	roomId, err := strconv.Atoi(input.RoomId)
	if err != nil {
		// Handle the error if the string cannot be converted
		fmt.Println("Error:", err)
		return
	}
	userId, err := strconv.Atoi(input.UserId)
	if err != nil {
		// Handle the error if the string cannot be converted
		fmt.Println("Error:", err)
		return
	}

	room := &model.Room{ID: roomId}
	user := model.User{ID: userId}

	conn, _ := database.ConnectDB()
	db, _ := conn.DB()
	defer db.Close()

	conn.First(&model.Room{ID: roomId})
	conn.Model(&room).Association("Users").Append(&user)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"roomId": roomId,
		"userId": userId,
	})
}

func LeaveRoom(c *gin.Context) {
	fmt.Println("Leave Room")

	var input struct {
		RoomId string `json:"roomId"`
		UserId string `json:"userId"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		// Handle JSON parsing error
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	roomId, err := strconv.Atoi(input.RoomId)
	if err != nil {
		// Handle the error if the string cannot be converted
		fmt.Println("Error:", err)
		return
	}
	userId, err := strconv.Atoi(input.UserId)
	if err != nil {
		// Handle the error if the string cannot be converted
		fmt.Println("Error:", err)
		return
	}

	room := &model.Room{ID: roomId}
	user := model.User{ID: userId}

	conn, _ := database.ConnectDB()
	db, _ := conn.DB()
	defer db.Close()
	conn.First(&model.Room{ID: roomId})
	conn.Model(&room).Association("Users").Delete(&user)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"roomId": roomId,
		"userId": userId,
	})
}
