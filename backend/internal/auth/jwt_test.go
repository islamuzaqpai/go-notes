package auth_test

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/islamuzaqpai/notes-app/internal/auth"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGenerateJWT(t *testing.T) {
	jwtKey := "test_secret"
	userId := 123

	tokenString, err := auth.GenerateJWT(userId, jwtKey)

	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})

	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	claims, ok := token.Claims.(jwt.MapClaims)
	assert.True(t, ok)

	assert.Equal(t, float64(userId), claims["user_id"])

	exp := int64(claims["exp"].(float64))
	assert.Greater(t, exp, time.Now().Unix())
}
