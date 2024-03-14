package helpers

import (
	"context"

	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// HasElement checks if the given username exists in the specified collection's array field under the document with the provided uid
func HasElement(collName, username, uid, key string) (bool, error) {
	// Open the specified collection
	coll := database.OpenCollection(database.Client, collName)

	// Convert uid string to ObjectId
	objID, err := primitive.ObjectIDFromHex(uid)
	if err != nil {
		return false, err
	}

	// Define a filter to find the document with the provided _id and containing the provided username in the array field
	filter := bson.M{"_id": objID, key: username}

	// Count the number of documents that match the filter
	count, err := coll.CountDocuments(context.Background(), filter)
	if err != nil {
		return false, err
	}

	// Return true if count is greater than 0, indicating the username exists in the specified document
	return count > 0, nil
}
