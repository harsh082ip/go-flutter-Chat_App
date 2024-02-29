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
		log.Println("---------HERE 10-----------")
		if !ok {
			return
		}

		log.Println("---------HERE 11-----------")
		c.Conn.WriteJSON(message)
		log.Println("---------HERE 12-----------")
	}
}

func (c *Client) readMessage(hub *Hub) {
	defer func() {
		hub.Unregister <- c
		c.Conn.Close()
	}()

	for {

		_, m, err := c.Conn.ReadMessage()
		log.Println("---------HERE 13-----------")
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

		log.Println("---------HERE 14-----------")
		hub.Broadcast <- msg
		log.Println("---------HERE 15-----------")
	}
}
