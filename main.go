package main

import (
	"fmt"

	controller "github.com/BelCattaneo/go-chat/app/controller"
	db "github.com/BelCattaneo/go-chat/app/db"
	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("Starting connection to database")
	conn, _ := db.ConnectDB()
	db.CreateTables(conn)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
		fmt.Println("pong madafaka!!!")
	})
	r.POST("/user/new", controller.CreateUser)

	r.Run() // listen and serve on 0.0.0.0:8080

}
