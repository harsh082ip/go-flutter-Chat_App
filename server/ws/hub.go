package ws

import (
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Hub struct {
	Rooms      map[string]*Room
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[string]*Room),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message),
	}
}

type Room struct {
	ID           primitive.ObjectID `bson:"_id"`
	RoomID       string             `json:"id"`
	Clients      []*Client          `json:"clients"`
	Participants []string           `json:"participants"`
}

func (h *Hub) Run() {
	for {
		select {
		case cl := <-h.Register:
			room, err := GetRoom(cl.RoomID)
			if err != nil {
				log.Println("Error Fetching Room Details: ", err)
				continue
			}
			if room != nil {
				room.Clients = append(room.Clients, cl)
				err := UpdateRoom(room)
				if err != nil {
					log.Println("Error updating the room : ", err)
				}
			}
			AddWebSocketConnection(cl.ConnId, cl.Conn)

		case cl := <-h.Unregister:
			room, err := GetRoom(cl.RoomID)
			if err != nil {
				log.Println("Error Fetching the Room Details: ", err)
				continue
			}
			if room != nil {
				for i, client := range room.Clients {
					if client == cl {
						room.Clients = append(room.Clients[:i], room.Clients[i+1:]...)
						err := UpdateRoom(room)
						if err != nil {
							log.Println("Error in Updating Room: ", err)
						}
						break
					}
				}
			}
			RemoveWebSocketConnection(cl.ConnId)

		case m := <-h.Broadcast:
			room, err := GetRoom(m.RoomID)
			if err != nil {
				log.Println("Error Fetching Room Details: ", err)
				continue
			}
			if room != nil {
				clientMap := make(map[string]*Client)
				for _, client := range room.Clients {
					clientMap[client.Id] = client
				}
				log.Println("Broadcasting message to all clients in the room...")
				log.Println("Room Clients:")
				for _, client := range clientMap {
					log.Println(client.Id)
				}
				h.BroadcastMessage(m, clientMap)
			}
		}
	}
}

// BroadcastMessage sends a message to all clients in the provided map
func (h *Hub) BroadcastMessage(m *Message, clients map[string]*Client) {
	for _, client := range clients {
		// log.Println("Broad...")
		// log.Println(client.ConnId)
		// Check if the WebSocket connection is available
		if conn := GetwebSocketConnection(client.ConnId); conn != nil {
			// log.Println("Broad...")
			// Send message over WebSocket connection
			err := conn.WriteJSON(m)
			// log.Println("Broaddd...")
			if err != nil {
				log.Println("Error sending message to client:", err)
			}
		}
	}
}

// func (h *Hub) Run() {
// 	for {
// 		select {
// 		case cl := <-h.Register:
// 			log.Println("Registered", cl)
// 		case cl := <-h.Unregister:
// 			log.Println("Unregistered", cl)
// 		case m := <-h.Broadcast:
// 			log.Println("Message", m)
// 		}
// 	}
// }
