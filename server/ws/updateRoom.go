package ws

import (
	"context"

	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateRoom(room *Room) error {
	filter := bson.M{"_id": room.ID}

	update := bson.M{
		"$set": bson.M{
			"clients": getSerializableClients(room.Clients),
		},
	}

	collName := "Rooms"
	coll := database.OpenCollection(database.Client, collName)

	result, err := coll.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	if result.ModifiedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}

func getSerializableClients(clients []*Client) []interface{} {
	var serializableClients []interface{}
	for _, client := range clients {
		serializableClient := map[string]interface{}{
			"Id":       client.Id,
			"Username": client.Username,
			"ConnID":   client.ConnId,
		}
		serializableClients = append(serializableClients, serializableClient)
	}
	return serializableClients
}
