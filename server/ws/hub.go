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
	Clients      map[string]*Client `json:"clients"`
	Participants []string           `json:"participants"`
}

func (h *Hub) Run() {
	for {
		select {
		case cl := <-h.Register:
			log.Println("---------HERE 4-----------")
			room, err := GetRoom(cl.RoomID)
			log.Println("---------HERE 5-----------")
			if err != nil {
				log.Println("Error Fetching Room Details: ", err)
				continue
			}
			log.Println("---------HERE 6-----------")
			if room != nil {
				room.Clients[cl.Id] = cl
				err := UpdateRoom(room)
				if err != nil {
					log.Println("Error updating the room : ", err)
				}
			}
			log.Println("---------HERE 7-----------")

			// Add the WebSocket connection to the connectionsMap
			AddWebSocketConnection(cl.Id, cl.Conn)

		case cl := <-h.Unregister:
			room, err := GetRoom(cl.RoomID)
			if err != nil {
				log.Println("Error Fetching the Room Details: ", err)
				continue
			}

			if room != nil {
				if _, ok := room.Clients[cl.Id]; ok {
					if len(room.Clients) != 0 {
						h.Broadcast <- &Message{
							Content:  "user left the chat",
							RoomID:   cl.RoomID,
							Username: cl.Username,
						}
					}

					delete(room.Clients, cl.Id)

					err := UpdateRoom(room)
					if err != nil {
						log.Println("Error in Updating Room: ", err)
					}
				}
			}

			// Remove the WebSocket connection from the connectionsMap
			RemoveWebSocketConnection(cl.Id)

		case m := <-h.Broadcast:
			log.Println("---------HERE 8-----------")
			room, err := GetRoom(m.RoomID)
			if err != nil {
				log.Println("Error Fetching Room Details: ", err)
			}

			if room != nil {
				// Call BroadcastMessage function to send message to clients
				h.BroadcastMessage(m, room.Clients)
			}
		}
	}
}

// BroadcastMessage sends a message to all clients in the provided map
func (h *Hub) BroadcastMessage(m *Message, clients map[string]*Client) {
	for _, client := range clients {
		// Check if the WebSocket connection is available
		if conn := GetwebSocketConnection(client.ConnId); conn != nil {
			// Send message over WebSocket connection
			err := conn.WriteJSON(m)
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
