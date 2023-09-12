package model

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       int    `gorm:"unique"`
	Username string `gorm:"unique"`
}

type Room struct {
	gorm.Model
	ID       int    `gorm:"unique"`
	Roomname string `gorm:"unique"`
	Users    []User `gorm:"many2many:room_users;"`
}

func CreateUser(user User) error {

	fmt.Println("Inserting in DB from model.CreateUser")
	return nil
}
