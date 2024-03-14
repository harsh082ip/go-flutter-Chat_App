package usersController

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/database"
	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/helpers"
	authhelper "github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/helpers/authHelper"
	wshelper "github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/helpers/wsHelper"
	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/models"
	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/ws"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddUserToRecentChatsAndCreateRoom(c *gin.Context) {
	// Extract parameters from the request
	uid := c.Param("uid")                  // userid of the user using the app
	username := c.Query("client2username") // this is user whom we are adding  (Present Client)
	jwt_token := c.Query("jwtkey")
	client1username := c.Query("client1username") // this is the user who is using the app, same user whose id is sent
	client2username := username                   // user which we want to add

	log.Println(uid)
	log.Println(client1username)
	// TODO: CHECK IF USERNAME AND UID BELONGS TO THE SAME USER

	// Check if all necessary parameters are provided
	if uid != "" && username != "" && jwt_token != "" && client2username != "" && client1username != client2username {
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

		err = helpers.CheckIfUidAndUsernameOfSameUser(uid, client1username)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "Bad Request",
				"error":  err.Error(),
			})
			return
		}

		client1usernameExists, _ := helpers.CheckIfDocumentExists(client1username, usersCollName, false)
		if !client1usernameExists {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "Invalid Client1 username, there is no such username",
			})
			return
		}

		// Check if document exists in Recent_Chats collection
		docExists, _ := helpers.CheckIfDocumentExists(uid, "Recent_Chats", true)

		// Open collection
		collName := "Recent_Chats"
		coll := database.OpenCollection(database.Client, collName)

		hub := &ws.Hub{}

		handler := ws.NewHandler(hub)
		var Room ws.Room

		// If all conditions met, update usernames if not present
		if usernameExists && UidExists && client1usernameExists && docExists {
			status, _ := helpers.HasElement(collName, username, uid, "usernames")
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

				roomStatus, err := wshelper.CheckIfRoomExists(client1username, client2username)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"status": "Error in looking for rooms",
						"error":  err.Error(),
					})
					return
				}

				if !roomStatus {

					room, status, err := handler.CreateRoom(client1username, client2username)
					if err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{
							"status": status,
							"error":  err,
						})
						return
					}

					Room = room
				}

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
					"room":   Room,
				})
				return
			} else {

				roomStatus, err := wshelper.CheckIfRoomExists(client1username, client2username)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"status": "Error looking for rooms",
						"error":  err.Error(),
					})
					return
				}

				if !roomStatus {

					room, status, err := handler.CreateRoom(client1username, client2username)
					if err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{
							"status": status,
							"error":  err,
						})
						return
					}

					c.JSON(http.StatusOK, gin.H{
						"status": "Username Already Present in Recent_Chats & Room Created",
						"room":   room,
					})
					return
				}

				c.JSON(http.StatusOK, gin.H{
					"status": "Usernames & Room already present",
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

		roomStatus, err := wshelper.CheckIfRoomExists(client1username, client2username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "Error in looking for rooms",
				"error":  err.Error(),
			})
			return
		}

		if !roomStatus {

			room, status, err := handler.CreateRoom(client1username, client2username)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status": status,
					"error":  err,
				})
				return
			}

			Room = room
		}

		_, err = coll.InsertOne(context.TODO(), user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "Error in Adding the Username to Recent Chats",
				"error":  err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "User Added & Room Created Successfully",
			"room":   Room,
		})
		return
	} else {
		// If necessary details are missing, return bad request
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Details Missing Or both usernames are equal",
			"error":  "please Send all the Information Or Correct it",
		})
	}
}
