package model

import "fmt"

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
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
