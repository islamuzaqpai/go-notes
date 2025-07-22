package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtKey = []byte("Up#ZDZP7ecml9Ff0cXbA")

func GenerateJWT(userId int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
