package usersController

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/database"
	authhelper "github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/helpers/authHelper"
	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

/*
GetUserByUsername retrieves user information by username from the database.
It requires a valid JWT token for authentication.

Parameters:
- c: Context object provided by Gin.
*/
func GetUserByUsername(c *gin.Context) {
	username := c.Param("username") // Extracting username from URL parameters
	jwt_token := c.Query("jwtkey")  // Extracting JWT token from query parameters

	_, err := authhelper.CheckAuthorization(jwt_token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "Unauthorized Access",
			"error":  err.Error(),
		})
		return
	}

	// Accessing MongoDB collection for user data
	collName := "Users"
	coll := database.OpenCollection(database.Client, collName)

	var result models.User
	// Querying the database for user information by username
	err = coll.FindOne(context.TODO(), bson.D{{"username", username}}).Decode(&result)

	if err != nil {
		// Handling errors during database query
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{
				"status": "Error in Document",
				"error":  "No User found with the given details",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Internal Server Error",
			"error":  "Error is searching for user",
		})
		return
	}

	// Returning user information on successful retrieval
	c.JSON(http.StatusOK, result)

}
