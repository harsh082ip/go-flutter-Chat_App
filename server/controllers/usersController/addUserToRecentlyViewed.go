package usersController

import (
	"context"
	// "encoding/json"
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/database"
	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/helpers"
	authhelper "github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/helpers/authHelper"
	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AddUserToRecentlyViewed is a handler function to add a user to the Recently Viewed collection
func AddUserToRecentlyViewed(c *gin.Context) {

	// Extract parameters from the request URL
	uid := c.Param("uid")
	userid := c.Query("userID")
	jwt_key := c.Query("jwt_key")

	// Check if all required parameters are provided
	if uid != "" && userid != "" && jwt_key != "" {

		// Check authorization using JWT token
		_, err := authhelper.CheckAuthorization(jwt_key)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "Unauthorized Access",
				"error":  err.Error(),
			})
			return
		}

		// Check if the document exists for the provided UID
		docExists, _ := helpers.CheckIfDocumentExists(uid)

		// Define MongoDB collection and model
		var user models.RecentlyViewed
		collName := "Recently_Viewed"
		coll := database.OpenCollection(database.Client, collName)

		if docExists {
			// Check if the user ID already exists in the document
			status, _ := helpers.HasElement(collName, userid)
			if !status {
				// If user ID does not exist, update the document by adding the user ID
				objId, err := primitive.ObjectIDFromHex(uid)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"status": "Error in creating object Id",
						"error":  err.Error(),
					})
					return
				}
				idFilter := bson.M{"_id": objId}
				update := bson.M{"$push": bson.M{"userids": userid}}

				_, err = coll.UpdateOne(context.Background(), idFilter, update)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"status": "Error in updating",
						"error":  err.Error(),
					})
					return
				}

				c.JSON(http.StatusOK, gin.H{
					"status": "Updated Successfully",
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"status": "next",
			})
			return
		}

		// If the document does not exist, create a new one
		user.ID, err = primitive.ObjectIDFromHex(uid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "Internal Server Error",
				"error":  err.Error(),
			})
			return
		}
		user.UserIDs = []string{userid}

		_, err = coll.InsertOne(context.TODO(), user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "Error in Inserting the doc",
				"error":  err.Error(),
			})
			return
		}
		// j, err := json.Marshal(user)
		c.JSON(http.StatusOK, gin.H{
			"status": "User Added to Recently Viewed",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Bad Request",
			"error":  "Incomplete or missing details",
		})
	}
}
