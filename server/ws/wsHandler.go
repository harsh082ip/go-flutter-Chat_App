package ws

import (
	"context"

	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/database"
	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Handler struct {
	hub *models.Hub
}

func NewHandler(h *models.Hub) *Handler {
	return &Handler{
		hub: h,
	}
}

func (h *Handler) CreateRoom(username1, username2 string) (models.Room, string, error) {

	id := primitive.NewObjectID()
	room := models.Room{
		ID:           id,
		RoomID:       id.Hex(),
		Clients:      make(map[string]*models.Client),
		Participants: []string{username1, username2},
	}

	collName := "Rooms"
	coll := database.OpenCollection(database.Client, collName)

	_, err := coll.InsertOne(context.Background(), room)
	if err != nil {
		return models.Room{}, "Cannot Create Room", err
	}

	return room, "Successfully Created Room", nil

}
