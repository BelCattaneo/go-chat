package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	model "github.com/BelCattaneo/go-chat/app/model"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("creating user!")
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)

	var user model.User
	err := decoder.Decode(&user)
	fmt.Println(user)

	if err != nil {
		fmt.Println(err)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err = model.CreateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
