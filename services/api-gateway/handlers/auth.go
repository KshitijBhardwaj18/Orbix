package handlers

import (
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"github.com/KshitijBhardwaj18/Orbix/services/api-gateway/utils"
    "github.com/KshitijBhardwaj18/Orbix/shared/models"
	"log"
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
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required,min=8"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return 
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal Server error"})
		return
	}

	user := &models.User{
		Email: req.Email,
		Username: req.Username,
		PasswordHash: hashedPassword,
	}

	if err := h.db.Create(&user).Error; err != nil {
		log.Printf("Database error creating user: %v", err)
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(201, gin.H{
		"message": "User registered successfully",
		"user_id": user.ID,
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req struct {
		Email string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid User Credentials "})
		log.Printf("Invalid User Credentials : %v", err)
		return
	}

	var user models.User

	if err := h.db.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(401, gin.H{"error": "Invalid User Credentials"})
		return
	}

	if !utils.VerifyPassword( req.Password, user.PasswordHash){
		c.JSON(401, gin.H{"error": "Invalid User Credentials"})
		return
	}

	token, err := utils.GenerateJWT(user.ID.String(), user.Email)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Login successful",
		"token":  token,
		"user": gin.H{
			"id":       user.ID,
            "email":    user.Email,
            "username": user.Username,
		},
	})
	
}



