package ws

import (
	"context"
	"log"

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

// Helper function to get serializable client information
func getSerializableClients(clients map[string]*Client) map[string]interface{} {
	serializableClients := make(map[string]interface{})
	for id, client := range clients {
		log.Println("Conn: ", client.Conn)
		serializableClient := map[string]interface{}{

			"Id":       client.Id,
			"Username": client.Username,
			"ConnID":   client.ConnId,
			// "Conn":     client.Conn,
			// Add other serializable fields from the client struct here
			// Exclude non-serializable fields like channels
		}
		serializableClients[id] = serializableClient
	}
	return serializableClients
}
