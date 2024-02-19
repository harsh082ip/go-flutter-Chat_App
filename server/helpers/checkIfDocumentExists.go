package helpers

import (
	"context"

	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CheckIfDocumentExists checks if a document exists in the specified collection based on the user ID or username
func CheckIfDocumentExists(user, collName string, isId bool) (bool, error) {

	// Open the specified collection
	coll := database.OpenCollection(database.Client, collName)

	// Check if user ID is provided
	if isId {
		// Convert user ID to ObjectID
		objId, err := primitive.ObjectIDFromHex(user)
		if err != nil {
			return false, err
		}

		// Define the filter to find documents by user ID
		userIDFilter := bson.D{{"_id", objId}}

		// Count the number of documents that match the filter
		count, err := coll.CountDocuments(context.TODO(), userIDFilter)
		if err != nil {
			return false, err
		}

		// If count is 0, document does not exist
		if count == 0 {
			return false, nil
		}

		// If count is not 0, document exists
		return true, nil
	}

	// If username is provided
	// Define the filter to find documents by username
	usernameFilter := bson.D{{"username", user}}

	// Count the number of documents that match the filter
	count, err := coll.CountDocuments(context.TODO(), usernameFilter)
	if err != nil {
		return false, err
	}

	// If count is 0, document does not exist
	if count == 0 {
		return false, nil
	}

	// If count is not 0, document exists
	return true, nil
}
