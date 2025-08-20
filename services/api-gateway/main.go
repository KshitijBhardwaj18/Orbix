package main

import (
	"log"

	"github.com/KshitijBhardwaj18/Orbix/services/api-gateway/handlers"
	"github.com/gin-gonic/gin"
	"github.com/KshitijBhardwaj18/Orbix/services/api-gateway/middleware"
	"github.com/KshitijBhardwaj18/Orbix/services/api-gateway/config"
)

func main() {
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	

	router := gin.Default()

	public := router.Group("/api/v1")
	authHandler := handlers.NewAuthHandler(db)
	{
		 public.POST("/auth/register", authHandler.Register)
		 public.POST("/auth/login", authHandler.Login)
		// public.GET("/health", healthHandler)
	}

	protected := router.Group("/api/v1")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/profile", )
	}

	log.Println("API Gateway is running on port :8080")
	router.Run(":8080")
}