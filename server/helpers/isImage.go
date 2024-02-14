package helpers

import (
	"mime/multipart"
	"net/http"
)

func IsImage(header *multipart.FileHeader) bool {
	// Open uploaded file
	file, err := header.Open()
	if err != nil {
		return false
	}
	defer file.Close()

	// Read the first 512 bytes to detect the file type
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return false
	}

	// Compare the file signature against known image signatures
	// Here we're checking for commonly used image file signatures: JPEG, PNG, GIF, BMP, and JPG
	contentType := http.DetectContentType(buffer)
	if contentType == "image/jpeg" ||
		contentType == "image/png" ||
		contentType == "image/gif" ||
		contentType == "image/bmp" ||
		contentType == "image/jpg" {
		return true
	}

	return false
}
