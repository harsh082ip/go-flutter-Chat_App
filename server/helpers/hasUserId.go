package helpers

import (
	"context"

	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/database"
	"go.mongodb.org/mongo-driver/bson"
)

// HasElement checks if the given user ID exists in the specified collection's array field
func HasElement(collName, userId string) (bool, error) {
	// Open the specified collection
	coll := database.OpenCollection(database.Client, collName)

	// Define a filter to find documents containing the provided user ID in the array field
	filter := bson.M{"userids": bson.M{"$elemMatch": bson.M{"$eq": userId}}}

	// Count the number of documents that match the filter
	count, err := coll.CountDocuments(context.Background(), filter)
	if err != nil {
		return false, err
	}

	// Return true if count is greater than 0, indicating the user ID exists in at least one document
	return count > 0, nil
}
