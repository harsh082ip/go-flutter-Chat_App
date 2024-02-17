package authhelper

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/models"
	"github.com/joho/godotenv"
)

// CheckAuthorization checks the authorization of a JWT token
func CheckAuthorization(jwt_token string) (string, error) {
	// Load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		return "", err
	}

	// Get JWT_SECRET_KEY from environment variables
	JWT_SECRET_KEY := os.Getenv("JWT_SECRET_KEY")

	// Create claims object to parse JWT token into
	claims := &models.Claims{}

	// Verify JWT token
	if JWT_SECRET_KEY != "" {
		// Parse JWT token with claims
		tkn, err := jwt.ParseWithClaims(string(jwt_token), claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(JWT_SECRET_KEY), nil
		})

		// Handle parsing errors
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				return "", err
			}
			return "", err
		}

		// Check if token is valid
		if !tkn.Valid {
			return "", err
		}

		// Return success if token is valid
		return "Success", nil
	}

	// Return error if JWT_SECRET_KEY is missing
	return "", fmt.Errorf("WARNING: SECRET KEY MISSING :/")
}
