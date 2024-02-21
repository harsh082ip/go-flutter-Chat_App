package usersController

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/database"
	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/helpers"
	authhelper "github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/helpers/authHelper"
	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddUserToRecentChats(c *gin.Context) {
	// Extract parameters from the request
	uid := c.Param("uid")
	username := c.Query("username")
	jwt_token := c.Query("jwtkey")

	// Check if all necessary parameters are provided
	if uid != "" && username != "" && jwt_token != "" {
		// Verify JWT token for authentication
		_, err := authhelper.CheckAuthorization(jwt_token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "Unauthorized Access",
				"error":  err.Error(),
			})
			return
		}

		// Check if user ID exists
		usersCollName := "Users"
		UidExists, _ := helpers.CheckIfDocumentExists(uid, usersCollName, true)
		if !UidExists {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "Invalid User ID, there is no such user",
			})
			return
		}

		// Check if username exists
		usernameExists, _ := helpers.CheckIfDocumentExists(username, usersCollName, false)
		if !usernameExists {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "Invalid Username, there is no such username present",
			})
			return
		}

		// Check if document exists in Recent_Chats collection
		docExists, _ := helpers.CheckIfDocumentExists(uid, "Recent_Chats", true)

		// Open collection
		collName := "Recent_Chats"
		coll := database.OpenCollection(database.Client, collName)

		// If all conditions met, update usernames if not present
		if usernameExists && UidExists && docExists {
			status, _ := helpers.HasElement(collName, username, "usernames")
			if !status {
				objID, err := primitive.ObjectIDFromHex(uid)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"status": "Error in Creating the ObjectID",
						"error":  err.Error(),
					})
					return
				}
				idFilter := bson.M{"_id": objID}
				update := bson.M{"$push": bson.M{"usernames": username}}
				_, err = coll.UpdateOne(context.Background(), idFilter, update)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"status": "Error in updating the usernames",
						"error":  err.Error(),
					})
					return
				}
				c.JSON(http.StatusOK, gin.H{
					"status": "Updated Successfully",
				})
				return
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"status": "Username already present",
				})
				return
			}
		}

		// If Recent_Chats document doesn't exist, insert new document
		var user models.RecentChats
		user.ID, err = primitive.ObjectIDFromHex(uid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "Error in Creating the Object ID",
				"error":  err.Error(),
			})
			return
		}
		user.Usernames = []string{username}
		_, err = coll.InsertOne(context.TODO(), user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "Error in Adding the Username to Recent Chats",
				"error":  err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "User Added Successfully",
		})
		return
	} else {
		// If necessary details are missing, return bad request
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Details Missing",
			"error":  "please Send all the Information",
		})
	}
}
