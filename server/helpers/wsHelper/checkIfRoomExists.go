package wshelper

import (
	"context"

	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/database"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckIfRoomExists(username1, username2 string) (bool, error) {

	collName := "Rooms"
	coll := database.OpenCollection(database.Client, collName)

	filter := bson.M{"participants": bson.M{"$all": []string{username1, username2}}}

	count, err := coll.CountDocuments(context.Background(), filter)
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	} else {
		return false, nil
	}
}
