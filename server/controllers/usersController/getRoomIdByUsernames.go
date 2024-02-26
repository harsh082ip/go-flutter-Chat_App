package usersController

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/database"
	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/helpers"
	authhelper "github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/helpers/authHelper"
	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/ws"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetRoomIDByUsernames(c *gin.Context) {

	username1 := c.Query("username1")
	username2 := c.Query("username2")
	jwt_token := c.Query("jwtkey")

	if username1 != "" && username2 != "" && jwt_token != "" {

		if username1 != username2 {

			_, err := authhelper.CheckAuthorization(jwt_token)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status": "Unauthorized Access",
					"error":  err.Error(),
				})
				return
			}

			usersCollName := "Users"

			username1Exists, _ := helpers.CheckIfDocumentExists(username1, usersCollName, false)
			if !username1Exists {
				c.JSON(http.StatusBadRequest, gin.H{
					"status": username1 + " does not exists",
					"error":  "Please Check username before going forword :/",
				})
				return
			}

			username2Exists, _ := helpers.CheckIfDocumentExists(username2, usersCollName, false)
			if !username2Exists {
				c.JSON(http.StatusBadRequest, gin.H{
					"status": username2 + " does not exists",
					"error":  "Please Check username before going forword :/",
				})
				return
			}

			roomCollName := "Rooms"
			coll := database.OpenCollection(database.Client, roomCollName)

			filter := bson.M{"participants": bson.M{"$all": []string{username1, username2}}}
			var room ws.Room
			err = coll.FindOne(context.Background(), filter).Decode(&room)
			if err != nil {
				if err == mongo.ErrNoDocuments {
					c.JSON(http.StatusBadRequest, gin.H{
						"status": "No Room Exists for the given usernames",
						"error":  err.Error(),
					})
					return
				}
				c.JSON(http.StatusInternalServerError, gin.H{
					"status": "Error in Looking for document",
					"error":  err.Error(),
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"status": "All Good!",
				"roomID": room.RoomID,
			})
			return

		}

		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Bad Request",
			"error":  "usernames can't be same",
		})
		return

	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Details Missing",
			"error":  "Please Provide all the details to continue",
		})
		return
	}
}
