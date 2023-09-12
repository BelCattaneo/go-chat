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
	/*
		query := `insert into posts(title, content) values($1, $2);`

		_, err := db.Exec(query, post.Title, post.Content)

		if err != nil {
			return err
		}
	*/

	fmt.Println("Inserting in DB from model.CreateUser")
	return nil
}
