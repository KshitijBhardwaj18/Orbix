package handlers

import (
	"errors"
	"log"

	"github.com/KshitijBhardwaj18/Orbix/shared/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserHandler struct {
	db *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{db:db}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	userIDstr := c.GetString("user_id")
	userID, err := uuid.Parse(userIDstr)
	if err != nil {
		log.Printf("invalid user_id in context: %v", err)
		c.JSON(401, gin.H{"error": "Invalid token"})
		return
	}

	var user models.User
	if err := h.db.Where("id = ?", userID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{"error": "User not found"})
			return
		}
		log.Printf("database error fetching user %s: %v", userID.String(), err)
		c.JSON(500, gin.H{"error": "Database error"})
		return
	}

	c.JSON(200, gin.H{
		"id":        user.ID,
		"email":     user.Email,
		"username":  user.Username,
		"created_at": user.CreatedAt,
	})
}
