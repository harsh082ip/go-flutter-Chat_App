package controllers

import (
	"context"
	"encoding/json"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/database"
	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/helpers"
	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SignUp handles the sign-up functionality.
func SignUp(c *gin.Context) {

	// Define collection name
	collName := "Users"
	coll := database.OpenCollection(database.DBinstance(), collName)

	var jsonData models.User

	// Parse the multipart form data
	err := c.Request.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form data"})
		return
	}

	// Get the JSON data from form
	jsonFile := c.Request.FormValue("data")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get JSON data"})
		return
	}

	// Decode JSON data into struct
	err = json.Unmarshal([]byte(jsonFile), &jsonData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode JSON data"})
		return
	}

	if jsonData.Name != "" && jsonData.Email != "" && jsonData.Password != "" && jsonData.Username != "" {
		// Check if email exists
		emailFilter := bson.D{{"email", jsonData.Email}}
		count, err := coll.CountDocuments(context.TODO(), emailFilter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "Error While Checking for Doc",
				"error":  err.Error(),
				"count":  count,
			})
			return
		}
		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "Doc Duplication not allowed",
				"error":  "this email already exists",
			})
			return
		}

		// Check if username exists
		usernameFilter := bson.D{{"username", jsonData.Username}}
		count, err = coll.CountDocuments(context.TODO(), usernameFilter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "Error While Checking for Doc",
				"error":  err.Error(),
			})
			return
		}
		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "Doc Duplication not allowed",
				"error":  "this username already exists",
			})
			return
		}

		// Generate ID for user
		jsonData.ID = primitive.NewObjectID()
		jsonData.UserId = jsonData.ID.Hex()

		// Hash password
		password, err := helpers.HashPassword(jsonData.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "Hashing Failed",
				"Error":  err.Error(),
			})
			return
		}

		// Check password validity
		if !helpers.CheckPasswordValidity(jsonData.Password) {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error":  "password constraints not fulfilled",
				"detail": "Password does not meet the required criteria: it must be at least 8 characters long, contain at least one lowercase letter, one uppercase letter, one special character, and one numeric digit.",
			})
			return
		}
		jsonData.Password = password

		// Verify email
		res, err := helpers.VerifyEmail(jsonData.Email)
		if !res {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "Email Verification Failed",
				"error":  err.Error(),
			})
			return
		}

		// Get profile picture from form
		file, header, err := c.Request.FormFile("profile_picture")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "Image Error",
				"error":  "Profile picture not found in form data",
			})
			return
		}

		// Check if the uploaded file is an image
		if !helpers.IsImage(header) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Uploaded file is not an image"})
			return
		}

		// Upload profile picture to MongoDB
		downloadURL, err := helpers.UploadToMongoDB(file, header)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "Image Upload Error",
				"error":  err.Error(),
			})
			return
		}
		jsonData.Profile_Pic_Url = downloadURL

		// Insert user data into database
		_, err = coll.InsertOne(context.TODO(), jsonData)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "Error is attempting to SignUp",
				"error":  err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "User SignUp Successful",
		})
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Missing Details",
			"error":  "Please fill the required fields",
		})
		return
	}
}
