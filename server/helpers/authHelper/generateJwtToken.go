package authhelper

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/models"
	"github.com/joho/godotenv"
)

func GenerateJwtToken(Useremail string) (string, error) {

	err := godotenv.Load(".env")
	if err != nil {
		return "", err
	}

	JWT_SECRET_KEY := os.Getenv("JWT_SECRET_KEY")

	if JWT_SECRET_KEY != "" {
		expirationTime := time.Now().Add(time.Hour * 24)

		claims := &models.Claims{
			CompanyName: "CyberSec Symposium",
			Email:       Useremail,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		key := []byte(JWT_SECRET_KEY)
		tokenString, err := token.SignedString(key)
		if err != nil {
			return "", err
		}

		return tokenString, nil
	}

	return "", fmt.Errorf("WARNING: SECRET KEY MISSING :/")

}
