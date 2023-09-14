package main

import (
	"fmt"

	"github.com/BelCattaneo/go-chat/app/controller"
	messagecontroller "github.com/BelCattaneo/go-chat/app/controller/message_controller"
	"github.com/BelCattaneo/go-chat/app/database"
	"github.com/gin-gonic/gin"
)

func initHandlers() {
	r := gin.Default()
	r.Static("/public", "./app/static")
	r.GET("/user", controller.GetUser)
	r.GET("/room", controller.GetRoom)
	r.POST("/user/new", controller.CreateUser)
	r.POST("/room/new", controller.CreateRoom)
	r.PUT("/user/room/enter", controller.EnterRoom)
	r.PUT("/user/room/leave", controller.LeaveRoom)

	r.GET("/upgrade", messagecontroller.UpgradeWithChannels)

	r.Run() // listen and serve on 0.0.0.0:8080
}

func main() {
	fmt.Println("Starting connection to database")

	database.SetupDB()

	initHandlers()

}
