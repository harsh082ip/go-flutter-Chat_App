package routes

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MiscRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/download/:id", func(c *gin.Context) {

		db := database.Client.Database("lpu_chathub")
		opts := options.GridFSBucket().SetName("Profile Pics")
		bucket, err := gridfs.NewBucket(db, opts)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}
		id := c.Param("id")

		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image ID"})
			return
		}

		stream, err := bucket.OpenDownloadStream(objID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open image"})
			return
		}
		defer stream.Close()

		c.Writer.Header().Set("Content-Type", "image/png")
		_, err = io.Copy(c.Writer, stream)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send image"})
			return
		}
	})
}
