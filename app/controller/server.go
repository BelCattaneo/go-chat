package controller

import (
	"fmt"

	controller "github.com/BelCattaneo/go-app/controller/chat"
	"github.com/gorilla/mux"
)

var router *mux.Router

func initHandlers() {
	router.HandleFunc("/api/user/new", controller.CreateUser).Methods("POST")
}

func Start() {
	router = mux.NewRouter()
	initHandlers()
	fmt.Printf("router initialized and listening on 3200\n")

}
