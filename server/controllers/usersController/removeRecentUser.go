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

func RemoveRecentUser(c *gin.Context) {

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

		// var checkUsers models.RecentlyViewed

		collName := "Recently_Viewed"

		docExists, err := helpers.CheckIfDocumentExists(uid, collName, true)
		if !docExists {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "Given uid does not Exists",
				"error":  err.Error(),
			})
			return
		}

		// TODO: Do it from here
		objID, err := primitive.ObjectIDFromHex(uid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "Error in Creating the object ID",
				"error":  err.Error(),
			})
			return
		}

		coll := database.OpenCollection(database.Client, collName)

		filter := bson.M{"_id": objID}
		update := bson.M{"$pull": bson.M{"userids": username}}
		result, err := coll.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "Error in updating the doc",
				"error":  err.Error(),
			})
			return
		}

		if result.ModifiedCount == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "Bad Request",
				"error":  "Username not found or Already Removed",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "User Removed Successfully",
		})

	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Details Missing",
			"error":  "Please Provide all the details",
		})
		return
	}
}
