package messagecontroller

import (
	"fmt"
	"strconv"

	"github.com/BelCattaneo/go-chat/app/controller/chat"
	"github.com/BelCattaneo/go-chat/app/controller/client"
	"github.com/BelCattaneo/go-chat/app/model"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var roomIdChatMap = chat.RoomIdChatMap{}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Upgrade(c *gin.Context) {
	fmt.Println("upgrade!")
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Error while upgrading")
		return
	}
	defer conn.Close()
	for {
		conn.WriteMessage(websocket.TextMessage, []byte("Hello, WebSocket!"))
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		if err := conn.WriteMessage(messageType, p); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(p))
	}
}

func UpgradeWithChannels(c *gin.Context) {
	fmt.Println("upgrade with channels!")
	srtUserId := c.DefaultQuery("user_id", "none")
	srtRoomId := c.DefaultQuery("room_id", "none")
	userId, err := strconv.Atoi(srtUserId)

	if err != nil {
		// Handle the error if the string cannot be converted
		fmt.Println("Error:", err)
		return
	}

	roomId, err := strconv.Atoi(srtRoomId)

	if err != nil {
		// Handle the error if the string cannot be converted
		fmt.Println("Error:", err)
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Error while upgrading")
		return
	}
	thisChat := roomIdChatMap.GetOrCreateChatForRoom(roomId)

	thisChat.Clients = append(thisChat.Clients, client.Client{
		User: model.User{ID: userId},
		Conn: conn,
	})

	defer conn.Close()
	fmt.Println("Connected to WS!")
	for {
		_, binaryMessage, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		thisChat.BroadcastMessage(binaryMessage)
	}
}

func sendValues(myIntChannel chan int) {

	for i := 0; i < 5; i++ {
		myIntChannel <- i //sending value
	}
	close(myIntChannel)
}

func MessageController() {
	myIntChannel := make(chan int)

	go sendValues(myIntChannel)

	for i := 0; i < 6; i++ {
		fmt.Println(<-myIntChannel) //receiving value
	}
}
