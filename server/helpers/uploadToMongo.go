package helpers

import (
	"fmt"
	"io"
	"mime/multipart"

	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// UploadToMongoDB uploads a file to MongoDB GridFS and returns the download URL.
func UploadToMongoDB(file multipart.File, header *multipart.FileHeader) (string, error) {

	// Get MongoDB database and options
	db := database.Client.Database("lpu_chathub")
	opts := options.GridFSBucket().SetName("Profile Pics")
	bucket, err := gridfs.NewBucket(db, opts)
	if err != nil {
		return "", err
	}

	// Open upload stream
	uploadStream, err := bucket.OpenUploadStream(header.Filename)
	if err != nil {
		return "", fmt.Errorf("Failed to open GridFS upload stream")
	}
	defer uploadStream.Close()

	// Copy file data to upload stream
	_, err = io.Copy(uploadStream, file)
	if err != nil {
		return "", err
	}

	// Get ObjectID and construct download URL
	objectID := uploadStream.FileID.(primitive.ObjectID)
	downloadURL := fmt.Sprintf("http://192.168.117.132:8006/download/%s", objectID.Hex())

	return downloadURL, nil
}
