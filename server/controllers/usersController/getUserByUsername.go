package usersController

import (
	"context"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/database"
	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/models"
	"github.com/joho/godotenv"
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

	claims := &models.Claims{} // Initializing JWT claims object

	// Loading environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		// Handle error if unable to load .env file
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Error Loading .env file",
		})
		return
	}

	JWT_SECRET_KEY := os.Getenv("JWT_SECRET_KEY") // Retrieving JWT secret key from environment variable

	if JWT_SECRET_KEY != "" {
		// Parsing JWT token with custom claims and verifying signature
		tkn, err := jwt.ParseWithClaims(string(jwt_token), claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(JWT_SECRET_KEY), nil
		})

		if err != nil {
			// Handling errors during JWT token parsing
			if err == jwt.ErrSignatureInvalid {
				c.JSON(http.StatusUnauthorized, gin.H{
					"status": "Unauthorized Access",
					"error":  err.Error(),
				})
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "Unauthorized Access",
				"error":  err.Error(),
			})
			return
		}

		if !tkn.Valid {
			// Handling invalid JWT token
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "Unauthorized Access",
				"error":  "Token is not Valid",
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

	} else {
		// Handling missing JWT secret key
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "WARNING: SECRET KEY MISSING :/",
		})
	}
}
