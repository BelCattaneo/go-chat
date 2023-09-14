package client

import (
	"github.com/BelCattaneo/go-chat/app/model"
	"github.com/gorilla/websocket"
)

type Client struct {
	User model.User
	Conn *websocket.Conn
}

func (client *Client) JoinChat() {
	// upgrade connection
	// create goroutine listening to chat goroutine messages and sending client messages
	// update chat struct with new user
}
