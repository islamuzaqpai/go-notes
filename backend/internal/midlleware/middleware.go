package midlleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"strings"
)

func AuthMiddleware(jwtKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid auth header"})
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid claims"})
			return
		}

		userId, ok := claims["user_id"].(float64)
		if !ok {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid user_id"})
			return
		}

		userIdInt := int(userId)
		c.Set("userID", userIdInt)

		c.Next()
	}
}
