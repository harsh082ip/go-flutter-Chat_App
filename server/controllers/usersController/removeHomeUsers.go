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

func RemoveHomeUsers(c *gin.Context) {

	uid := c.Param("uid")
	username := c.Query("username")
	jwt_token := c.Query("jwtkey")

	if uid != "" && username != "" && jwt_token != "" {

		_, err := authhelper.CheckAuthorization(jwt_token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "Unauthorized Access",
				"error":  err.Error(),
			})
			return
		}

		uidExists, _ := helpers.CheckIfDocumentExists(uid, "Users", true)
		if !uidExists {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "Uid not found",
				"error":  "There is no such user with the provided uid",
			})
			return
		}

		usernameExists, _ := helpers.CheckIfDocumentExists(username, "Users", false)
		if !usernameExists {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "Username not found",
				"error":  "Theere is no such user with the provided username",
			})
			return
		}

		docExists, _ := helpers.CheckIfDocumentExists(uid, "Recent_Chats", true)
		if !docExists {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "No users in Recent Chats",
				"error":  "Provided uid does not have any users in Recent Chats",
			})
			return
		}

		objID, err := primitive.ObjectIDFromHex(uid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "Error in Creating the objectID",
				"error":  err.Error(),
			})
			return
		}

		filter := bson.M{"_id": objID}
		update := bson.M{"$pull": bson.M{"usernames": username}}
		collName := "Recent_Chats"
		coll := database.OpenCollection(database.Client, collName)

		result, err := coll.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "Error in Updating the doc",
				"error":  err.Error(),
			})
			return
		}

		if result.ModifiedCount == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "Bad Request",
				"error":  "Username Not found or already removed",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "username removed successfully",
		})

		return

	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Missing Details",
			"error":  "Please Provide all the details",
		})
		return
	}

}
