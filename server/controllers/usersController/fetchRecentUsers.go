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

func FetchRecentUsers(c *gin.Context) {

	uid := c.Param("uid")
	jwt_token := c.Query("jwtkey")

	if uid != "" && jwt_token != "" {

		_, err := authhelper.CheckAuthorization(jwt_token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "Unauthorized Access",
				"error":  err.Error(),
			})
			return
		}

		collName := "Recently_Viewed"
		docExists, _ := helpers.CheckIfDocumentExists(uid, collName, true)
		if !docExists {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "User id is invalid",
				"error":  "User id does not have recent Viewed people",
			})
			return
		}

		objId, err := primitive.ObjectIDFromHex(uid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "Error in Creatinig an ObjectId",
				"error":  err.Error(),
			})
			return
		}
		filter := bson.D{{"_id", objId}}
		var recent_view models.RecentlyViewed

		coll := database.OpenCollection(database.Client, collName)

		err = coll.FindOne(context.TODO(), filter).Decode(&recent_view)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusBadRequest, gin.H{
					"status": "Doc. Not Found with the uid",
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

		usersCollName := "Users"
		usersColl := database.OpenCollection(database.Client, usersCollName)
		fmt.Println("User IDs to search:", recent_view.UserIDs) // Log user IDs for debugging
		cursor, err := usersColl.Find(context.TODO(), bson.M{"username": bson.M{"$in": recent_view.UserIDs}})
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

		var users []models.User
		if err := cursor.All(context.TODO(), &users); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "Error in decoding users",
				"error":  err.Error(),
			})
			return
		}

		if users != nil {
			c.JSON(http.StatusOK, users)
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "No Users Found",
			"error":  "Recents may be empty",
		})

	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Missing Details",
			"error":  "Please provide all the details",
		})
		return
	}
}
