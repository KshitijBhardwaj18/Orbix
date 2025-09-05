package middleware

import (
	
	"github.com/gin-gonic/gin"
	"log"
	"github.com/KshitijBhardwaj18/Orbix/services/api-gateway/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("authToken")

		if err != nil {
			c.JSON(401,gin.H{"error": "Authentication required"})
			c.Abort( )
		}

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