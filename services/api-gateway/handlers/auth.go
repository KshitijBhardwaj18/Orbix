package handlers

import (
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"github.com/KshitijBhardwaj18/Orbix/services/api-gateway/utils"
    "github.com/KshitijBhardwaj18/Orbix/shared/models" 
)

type AuthHandler struct {
	db *gorm.DB
}

func NewAuthHandler(db *gorm.DB) *AuthHandler {
	return &AuthHandler{db:db}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req struct {
		Email string `json:"email" binding:"required,email"`
		Username string `json:"username " binding:"required"`
		Password string `json:"password" binding:"required,min=8"`
	}

	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return 
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(500, g.H{"error": "Internal Server error"})
		return
	}

	user := &models.User{
		Email: req.Email,
		Username: req.Username,
		PasswordHash: hashedPassword,
	}

	if err := h.db.Create(user); err != nil {
		c.JSON(400, gin.H{"error": "User already exists"})
		return
	}

	c.JSON(201, gin.H{
		"message": "User registered successfully",
		"user_id": user.ID,
	})
}

