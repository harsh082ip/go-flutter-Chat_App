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
	"go.mongodb.org/mongo-driver/mongo"
)

func FetchChatsFromDB(c *gin.Context) {

	roomId := c.Param("roomId")
	jwt_token := c.Query("jwt_key")

	if roomId != "" && jwt_token != "" {

		_, err := authhelper.CheckAuthorization(jwt_token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "Unauthorized Access",
				"error":  err.Error(),
			})
			return
		}

		_, err = helpers.CheckIfDocumentExists(roomId, "Rooms", true)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "No Room Available with the given ID",
				"error":  err.Error(),
			})
			return
		}

		collName := "Chats"
		coll := database.OpenCollection(database.Client, collName)

		var conversation models.Conversation

		objID, err := primitive.ObjectIDFromHex(roomId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "Error Creating ObjectID",
				"error":  err.Error(),
			})
			return
		}

		filter := bson.M{"_id": objID}

		err = coll.FindOne(context.Background(), filter).Decode(&conversation)

		if err != nil {
			// Handling errors during database query
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusOK, gin.H{
					"messages": []string{},
				})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "Internal Server Error",
				"error":  "Error is searching for user",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"messages": conversation.Messages,
		})
		return

	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Missing Information",
			"error":  "Please Provide all the details to continue",
		})
		return
	}
}
