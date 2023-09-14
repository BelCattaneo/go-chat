package chat

import (
	"fmt"

	"github.com/BelCattaneo/go-chat/app/controller/client"
	"github.com/gorilla/websocket"
)

type RoomIdChatMap map[int]*Chat

func NewRoomIdChatMap() RoomIdChatMap {
	return RoomIdChatMap{}
}

func (chatMap RoomIdChatMap) GetChatClients(roomId int) []client.Client {
	return chatMap[roomId].Clients
}

func (chatMap RoomIdChatMap) GetOrCreateChatForRoom(roomId int) *Chat {
	if chatMap[roomId] == nil {
		chatMap[roomId] = &Chat{}
	}
	return chatMap[roomId]
}

func (chatMap RoomIdChatMap) GetChatForRoom(roomId int) *Chat {
	return chatMap[roomId]
}

type Chat struct {
	Clients []client.Client
}

func (chat *Chat) ReceiveMessages() {
	// when room is not empty creates ws connection with goroutine listening to messages
}

func (chat *Chat) BroadcastMessage(message []byte) {
	// when a message is received it sends the same message tho chat clients
	for _, client := range chat.Clients {
		fmt.Printf("sending message to user with ID: %+v\n", client.User.ID)
		client.Conn.WriteMessage(websocket.TextMessage, message)
	}
}
