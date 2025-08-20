package middleware

import (
	"strings"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(401, gin.H{"error": "Invalid token format "})
			c.Abort()
			return 
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		if token == "valid-token" {
			c.Set("user_id","user-123")
			c.Next()
		}else{
			c.JSON(401,gin.H{"error": "Invalid token"})
			c.Abort()
		}


	}
}