// Package usersController provides handlers for user-related endpoints.
package usersController

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/database"
	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/helpers"
	authhelper "github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/helpers/authHelper"
	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// FetchRecentUsers fetches recent users for a given user ID.
func FetchRecentUsers(c *gin.Context) {

	// Extract user ID and JWT token from request parameters.
	uid := c.Param("uid")
	jwtToken := c.Query("jwtkey")

	// Ensure both user ID and JWT token are provided.
	if uid != "" && jwtToken != "" {

		// Check the validity of the JWT token.
		_, err := authhelper.CheckAuthorization(jwtToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "Unauthorized Access",
				"error":  err.Error(),
			})
			return
		}

		// Check if a document exists for the recent user views.
		collName := "Recently_Viewed"
		docExists, _ := helpers.CheckIfDocumentExists(uid, collName, true)
		if !docExists {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "User id is invalid",
				"error":  "User id does not have recent Viewed people",
			})
			return
		}

		// Convert user ID string to ObjectID.
		objID, err := primitive.ObjectIDFromHex(uid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "Error in Creating an ObjectId",
				"error":  err.Error(),
			})
			return
		}
		filter := bson.D{{"_id", objID}}
		var recentView models.RecentlyViewed

		// Retrieve the recently viewed users document from the database.
		coll := database.OpenCollection(database.Client, collName)

		err = coll.FindOne(context.TODO(), filter).Decode(&recentView)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusBadRequest, gin.H{
					"status": "Document Not Found with the UID",
					"error":  err.Error(),
				})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "Error in Searching for the Users",
				"error":  err.Error(),
			})
			return
		}

		// Retrieve users based on their usernames.
		usersCollName := "Users"
		usersColl := database.OpenCollection(database.Client, usersCollName)
		fmt.Println("User IDs to search:", recentView.UserIDs) // Log user IDs for debugging
		cursor, err := usersColl.Find(context.TODO(), bson.M{"username": bson.M{"$in": recentView.UserIDs}})
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status": "No users found",
					"error":  err.Error(),
				})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "Error in fetching users",
				"error":  err.Error(),
			})
			return
		}

		// Decode the retrieved users.
		var users []models.User
		if err := cursor.All(context.TODO(), &users); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "Error in decoding users",
				"error":  err.Error(),
			})
			return
		}

		// Return the fetched users.
		if users != nil {
			c.JSON(http.StatusOK, users)
			return
		}

		// No users found.
		c.JSON(http.StatusOK, gin.H{
			"status": "No Users Found",
			"error":  "Recent views may be empty",
		})

	} else {
		// Missing user ID or JWT token.
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Missing Details",
			"error":  "Please provide all the details",
		})
		return
	}
}
