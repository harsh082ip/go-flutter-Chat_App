package usersController

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/database"
	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/helpers"
	authhelper "github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/helpers/authHelper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// RemoveRecentUser is a function to remove a specified username from the user's recent list.
func RemoveRecentUser(c *gin.Context) {

	// Extract uid, username, and jwt_token from the request.
	uid := c.Param("uid") // user using the app
	username := c.Query("username")
	jwt_token := c.Query("jwtkey")

	// Check if all required details are provided.
	if uid != "" && username != "" && jwt_token != "" {

		// Validate the JWT token to ensure the request is authorized.
		_, err := authhelper.CheckAuthorization(jwt_token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "Unauthorized Access",
				"error":  err.Error(),
			})
			return
		}

		// Define the collection name where the user's recent list is stored.
		collName := "Recently_Viewed"

		// Check if the document (user's recent list) exists.
		docExists, err := helpers.CheckIfDocumentExists(uid, collName, true)
		if !docExists {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "Given uid does not Exists",
				"error":  err.Error(),
			})
			return
		}

		// Convert the uid from string to MongoDB's ObjectID format.
		objID, err := primitive.ObjectIDFromHex(uid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "Error in Creating the object ID",
				"error":  err.Error(),
			})
			return
		}

		// Open the MongoDB collection.
		coll := database.OpenCollection(database.Client, collName)

		// Define the filter to locate the document and the update operation to remove the username.
		filter := bson.M{"_id": objID}
		update := bson.M{"$pull": bson.M{"userids": username}}

		// Perform the update operation.
		result, err := coll.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "Error in updating the doc",
				"error":  err.Error(),
			})
			return
		}

		// Check if the operation modified any document.
		if result.ModifiedCount == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "Bad Request",
				"error":  "Username not found or Already Removed",
			})
			return
		}

		// Confirm successful removal.
		c.JSON(http.StatusOK, gin.H{
			"status": "User Removed Successfully",
		})

	} else {
		// Respond with an error if any of the required details are missing.
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Details Missing",
			"error":  "Please Provide all the details",
		})
		return
	}
}
