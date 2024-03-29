package ws

import (
	"context"

	// "log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/database"
	wshelper "github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/helpers/wsHelper"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Handler struct {
	hub *Hub
}

func NewHandler(h *Hub) *Handler {
	return &Handler{
		hub: h,
	}
}

func (h *Handler) CreateRoom(username1, username2 string) (Room, string, error) {

	id := primitive.NewObjectID()
	room := Room{
		ID:           id,
		RoomID:       id.Hex(),
		Clients:      []*Client{},
		Participants: []string{username1, username2},
	}

	collName := "Rooms"
	coll := database.OpenCollection(database.Client, collName)

	_, err := coll.InsertOne(context.Background(), room)
	if err != nil {
		return Room{}, "Cannot Create Room", err
	}

	return room, "Successfully Created Room", nil

}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) JoinRoom(c *gin.Context) {

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Cannot Upgrade to ws connection",
			"error":  err.Error(),
		})
		return
	}

	roomID := c.Param("roomId")
	clientID := c.Query("uid")
	username := c.Query("username")

	connId := wshelper.GenerateConnID()

	cl := &Client{
		ConnId:   connId,
		Conn:     conn,
		Message:  make(chan *Message, 10),
		Id:       clientID,
		RoomID:   roomID,
		Username: username,
	}

	m := &Message{
		Content:  "A new user has joined the room",
		RoomID:   roomID,
		Username: username,
	}

	h.hub.Register <- cl
	h.hub.Broadcast <- m

	go cl.WriteMessage()
	go cl.readMessage(h.hub)
}

var ExportedHandler = &Handler{
	hub: NewHub(),
}
