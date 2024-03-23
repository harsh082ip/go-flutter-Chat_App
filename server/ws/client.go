package ws

import (
	"log"

	"github.com/gorilla/websocket"
)

// import "github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/models"

type Client struct {
	Conn     *websocket.Conn
	ConnId   string
	Message  chan *Message
	Id       string `json:"id"`
	RoomID   string `json:"roomId"`
	Username string `json:"username"`
}

type Message struct {
	Content  string `json:"content"`
	RoomID   string `json:"roomId"`
	Username string `json:"username"`
}

var ConnectionsMap map[string]*websocket.Conn = make(map[string]*websocket.Conn)

func (c *Client) WriteMessage() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		message, ok := <-c.Message
		if !ok {
			return
		}

		c.Conn.WriteJSON(message)
	}
}

func (c *Client) readMessage(hub *Hub) {
	defer func() {
		log.Println("ConnIddd:", c.ConnId)
		log.Println("Unregistered...")
		hub.Unregister <- c
		log.Println("Unregistered 2...")
		c.Conn.Close()
		log.Println("Unregistered 3...")
	}()

	for {
		_, m, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		msg := &Message{
			Content:  string(m),
			RoomID:   c.RoomID,
			Username: c.Username,
		}

		hub.Broadcast <- msg
		messageChannel <- msg
	}
}
