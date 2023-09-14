package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/BelCattaneo/go-chat/app/database"
	"github.com/BelCattaneo/go-chat/app/model"
)

func getRoomAndUserFromRequest(c *gin.Context) (model.Room, model.User) {
	var input struct {
		RoomId string `json:"roomId"`
		UserId string `json:"userId"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		// Handle JSON parsing error
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return model.Room{}, model.User{}
	}

	roomId, err := strconv.Atoi(input.RoomId)
	if err != nil {
		// Handle the error if the string cannot be converted
		fmt.Println("Error:", err)
		return model.Room{}, model.User{}
	}
	userId, err := strconv.Atoi(input.UserId)
	if err != nil {
		// Handle the error if the string cannot be converted
		fmt.Println("Error:", err)
		return model.Room{}, model.User{}
	}
	return model.Room{ID: roomId}, model.User{ID: userId}
}

func GetUser(c *gin.Context) {
	fmt.Println("Get User")
	srtUserId := c.DefaultQuery("user_id", "none")
	userId, err := strconv.Atoi(srtUserId)

	if err != nil {
		// Handle the error if the string cannot be converted
		fmt.Println("Error:", err)
		return
	}

	conn, _ := database.ConnectDB()
	db, _ := conn.DB()
	defer db.Close()
	result := map[string]interface{}{}
	conn.Model(&model.User{ID: userId}).First(&result)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"result": result,
	})
}

func GetRoom(c *gin.Context) {
	fmt.Println("Get Room")
	srtRoomId := c.DefaultQuery("room_id", "none")
	roomId, err := strconv.Atoi(srtRoomId)

	if err != nil {
		// Handle the error if the string cannot be converted
		fmt.Println("Error:", err)
		return
	}

	conn, _ := database.ConnectDB()
	db, _ := conn.DB()
	defer db.Close()

	var room model.Room
	if err := conn.Where("id = ?", roomId).Preload("Users").First(&room).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"result": room,
	})
}

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

	room, user := getRoomAndUserFromRequest(c)

	conn, _ := database.ConnectDB()
	db, _ := conn.DB()
	defer db.Close()

	conn.First(&model.Room{ID: room.ID})
	conn.Model(&room).Association("Users").Append(&user)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"roomId": room.ID,
		"userId": user.ID,
	})
}

func LeaveRoom(c *gin.Context) {
	fmt.Println("Leave Room")

	room, user := getRoomAndUserFromRequest(c)

	conn, _ := database.ConnectDB()
	db, _ := conn.DB()
	defer db.Close()
	conn.First(&model.Room{ID: room.ID})
	conn.Model(&room).Association("Users").Delete(&user)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"roomId": room.ID,
		"userId": user.ID,
	})
}
