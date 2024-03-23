package ws

import (
	"context"
	"log"

	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var messageChannel = make(chan *Message)

// Start listening for messages and saving them to the database
func SaveChatsToDatabase() {
	go func() {
		for receivedMsg := range messageChannel {

			err := SaveMessages(receivedMsg)
			if err != nil {
				log.Println(err)
			}
			log.Println("New Message in Room ID", receivedMsg.RoomID)
			continue
		}
	}()
}

func SaveMessages(message *Message) error {

	objId, err := primitive.ObjectIDFromHex(message.RoomID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objId}

	// Create a BSON representation of the message without the RoomID
	msg := bson.M{
		"content":  message.Content,
		"username": message.Username,
	}

	// Update operation to push the new message to the 'messages' array
	update := bson.M{
		"$push": bson.M{"messages": msg},
	}

	// Upsert option to insert the document if it doesn't exist
	opts := options.Update().SetUpsert(true)

	collName := "Chats"
	coll := database.OpenCollection(database.Client, collName)

	_, err = coll.UpdateOne(context.Background(), filter, update, opts)
	if err != nil {
		return err
	}

	return nil
}
