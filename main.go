package main

import (
	"fmt"

	"github.com/BelCattaneo/go-chat/app/controller"
	"github.com/BelCattaneo/go-chat/app/database"
	"github.com/gin-gonic/gin"
)

func initHandlers() {
	r := gin.Default()

	r.POST("/user/new", controller.CreateUser)

	r.Run() // listen and serve on 0.0.0.0:8080
}

func main() {

	fmt.Println("Starting connection to database")
	database.ConnectDB()
	database.CreateTables()

	initHandlers()

}
