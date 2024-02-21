package usersController

import (
	"context"
	"net/http"

	// Import necessary packages for HTTP handling, MongoDB operations, and custom utility functions.
	"github.com/gin-gonic/gin"
	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/database"
	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/helpers"
	authhelper "github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/helpers/authHelper"
	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func FetchHomeUsers(c *gin.Context) {
	// Retrieve user ID and JWT token from the request.
	uid := c.Param("uid")
	jwt_token := c.Query("jwtkey")

	// Check if both user ID and JWT token are provided.
	if uid != "" && jwt_token != "" {
		// Validate JWT token for authorization.
		_, err := authhelper.CheckAuthorization(jwt_token)
		if err != nil {
			// Return error response if token is invalid.
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "Unauthorized Access",
				"error":  err.Error(),
			})
			return
		}

		// Check if the user ID exists in the database.
		uidValid, _ := helpers.CheckIfDocumentExists(uid, "Users", true)
		if !uidValid {
			// Return error response if user ID is invalid.
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "User Id is Invalid",
				"error":  "There is no such user with the given uid",
			})
			return
		}

		// Check if the user has recent chats.
		collName := "Recent_Chats"
		docExists, _ := helpers.CheckIfDocumentExists(uid, collName, true)
		if !docExists {
			// Return error response if no recent chats are found.
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "No users Found in recent chats",
				"error":  "The User does not have recent users",
			})
			return
		}

		// Convert user ID from string to MongoDB's ObjectID.
		objId, err := primitive.ObjectIDFromHex(uid)
		if err != nil {
			// Return error if there's an issue with ObjectID conversion.
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "Error Creating the ObjectID",
				"error":  err.Error(),
			})
			return
		}

		// Query the Recent_Chats collection for the current user's chats.
		var recentChats models.RecentChats
		filter := bson.D{{"_id", objId}}
		coll := database.OpenCollection(database.Client, collName)
		err = coll.FindOne(context.TODO(), filter).Decode(&recentChats)
		if err != nil {
			// Handle errors in fetching or decoding the document.
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusBadRequest, gin.H{
					"status": "No Doc Found with the given uid",
					"error":  err.Error(),
				})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "Error is Searching for the provided uid",
				"error":  err.Error(),
			})
			return
		}

		// Fetch users from the Users collection based on usernames in recent chats.
		usersCollName := "Users"
		usersColl := database.OpenCollection(database.Client, usersCollName)
		cursor, err := usersColl.Find(context.TODO(), bson.M{"username": bson.M{"$in": recentChats.Usernames}})
		if err != nil {
			// Handle errors in fetching users.
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "Error in Fetching Users",
				"error":  err.Error(),
			})
			return
		}

		// Decode the cursor into a slice of User models.
		var users []models.User
		if err := cursor.All(context.TODO(), &users); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "Error in decoding users",
				"error":  err.Error(),
			})
			return
		}

		// Return the users if found.
		if users != nil {
			c.JSON(http.StatusOK, users)
			return
		}

		// Return error if no users are found.
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "No Users Found",
			"error":  "Recent views may be empty",
		})

	} else {
		// Return error if necessary details are missing.
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Details Missing",
			"error":  "Please provide all the details",
		})
		return
	}
}

/*

"username": bson.M{"$in": recentChats.Usernames}: This part of the filter specifies
 that you're looking for documents where the username field matches any value in the
 list recentChats.Usernames. The $in operator is used in MongoDB queries to select
 documents where the value of a field equals any value in the specified array.
 In this case, recentChats.Usernames is an array

*/
