package main

import (
	"log"

	"github.com/KshitijBhardwaj18/Orbix/services/api-gateway/handlers"
	"github.com/gin-gonic/gin"
	"github.com/KshitijBhardwaj18/Orbix/services/api-gateway/middleware"
	"github.com/KshitijBhardwaj18/Orbix/services/api-gateway/config"
)

func main() {
	// db, err := config.ConnectDatabase()
	// if err != nil {
	// 	log.Fatalf("Failed to connect to database: %v", err)
	// }

	

	router := gin.Default()

	// public := router.Group("/api/v1")

	{
		// public.POST("/auth/register", handlers.RegisterHandler)
		// public.POST("/auth/login", handlers.LoginHandler)
		// public.GET("/health", healthHandler)
	}

	// protected := router.Group("/api/v1")
	// protected.Use(authMiddleware())
	{
		// protected.GET("/auth/profile", profileHandler)
		// protected.POST("/auth/logout", logoutHandler)
	}

	log.Println("API Gateway is running on port :8080")
	router.Run(":8080")
}