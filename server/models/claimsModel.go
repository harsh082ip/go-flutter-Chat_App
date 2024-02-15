package models

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	CompanyName string `json:"comp_name"`
	Email       string `json:"email"`
	jwt.StandardClaims
}
