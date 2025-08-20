package middleware

import (
	"strings"
	"github.com/gin-gonic/gin"
	"log"
	"github.com/KshitijBhardwaj18/Orbix/services/api-gateway/utils"
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

		claims, err := utils.ValidateJWT(token)

		if err != nil {
			c.JSON(401, gin.H{"error": "Authentication Problem"})
			log.Printf("Authentication Problem : %v", err)
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("email", claims.Email)
		c.Next()
	}
}