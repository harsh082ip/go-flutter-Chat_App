package helpers

import (
	"context"
	"fmt"

	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CheckIfDocumentExists(userid string) (bool, error) {

	collName := "Recently_Viewed"
	coll := database.OpenCollection(database.Client, collName)

	objId, err := primitive.ObjectIDFromHex(userid)
	if err != nil {

		return false, err
	}
	userid_filter := bson.D{{"_id", objId}}

	count, err := coll.CountDocuments(context.TODO(), userid_filter)
	if err != nil {
		return false, err
	}
	fmt.Println(count)
	if count == 0 {
		return false, fmt.Errorf("doc. already exists")
	}

	return true, nil
}
