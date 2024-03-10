package helpers

import (
	"context"
	"fmt"

	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/database"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckIfUidAndUsernameOfSameUser(uid, username string) error {

	collName := "Users"
	coll := database.OpenCollection(database.Client, collName)

	filter := bson.D{
		{"userid", uid},
		{"username", username},
	}

	count, err := coll.CountDocuments(context.Background(), filter)
	if err != nil {
		return err
	}

	if count > 0 {
		return nil
	}
	return fmt.Errorf("none of the user found with the given uid and username")
}
