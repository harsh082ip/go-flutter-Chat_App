package ws

import (
	"context"
	"log"

	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetRoom(roomID string) (*Room, error) {
	log.Print("--------Here-----------")
	collName := "Rooms"
	coll := database.OpenCollection(database.Client, collName)

	filter := bson.D{{"roomid", roomID}}

	var room *Room

	err := coll.FindOne(context.Background(), filter).Decode(&room)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		return nil, err
	}

	return room, nil
}
