package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateJWT(userId int, jwtKey string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtKey))
}
