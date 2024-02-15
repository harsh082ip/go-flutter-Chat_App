package authhelper

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/models"
	"github.com/joho/godotenv"
)

func GenerateRefreshToken(c *gin.Context) {
	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		// Handle error if loading environment variables fails
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error Loading the .env file",
		})
		return
	}

	// Get JWT secret key from environment variables
	JWT_SECRET_KEY := os.Getenv("JWT_SECRET_KEY")

	if JWT_SECRET_KEY != "" {
		// Read request body
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			// Handle error if reading body fails
			c.String(http.StatusInternalServerError, "Error reading request body")
			return
		}
		defer c.Request.Body.Close() // Close the body to prevent leaks

		// Parse token string from request body
		tokenStr := body
		claims := &models.Claims{}

		// Parse JWT token with claims
		tkn, err := jwt.ParseWithClaims(string(tokenStr), claims,
			func(t *jwt.Token) (interface{}, error) {
				// Provide JWT secret key as interface
				return []byte(JWT_SECRET_KEY), nil
			})
		if err != nil {
			// Handle JWT parsing errors
			if err == jwt.ErrSignatureInvalid {
				c.JSON(http.StatusUnauthorized, gin.H{
					"status": "Unauthorized",
					"error":  err.Error(),
				})
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "Unauthorized",
				"error":  err.Error(),
			})
			return
		}

		if !tkn.Valid {
			// Handle invalid token
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "Unauthorized Access",
				"error":  "token is invalid",
			})
			return
		}

		if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
			// Handle case when token expiration is too far in future
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "Bad Request",
				"error":  "Cannot Generate Refresh token now",
			})
			return
		}

		// Set expiration time for refresh token
		expirationTime := time.Now().Add(time.Hour * 24)
		claims.ExpiresAt = expirationTime.Unix()

		// Generate new JWT token with updated claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		key := []byte(JWT_SECRET_KEY)
		tokenString, err := token.SignedString(key)

		if err != nil {
			// Handle error in generating JWT token
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "Error in generating Refresh Token",
			})
			return
		}

		// Return success response with new JWT token
		c.JSON(http.StatusOK, gin.H{
			"status":    "All Good!!",
			"Jwt_Token": tokenString,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "WARNING: SECRET KEY MISSING :/",
		})
	}
}
