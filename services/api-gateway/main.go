package main

import (
	"log"

	"github.com/KshitijBhardwaj18/Orbix/services/api-gateway/config"
	"github.com/KshitijBhardwaj18/Orbix/services/api-gateway/handlers"
	"github.com/KshitijBhardwaj18/Orbix/services/api-gateway/middleware"
	"github.com/KshitijBhardwaj18/Orbix/shared/broker"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	Broker := broker.NewRedisClient()

    authHandler := handlers.NewAuthHandler(db)
	orderHandler := handlers.NewOrderHandler(Broker)
	

	router := gin.Default()

	public := router.Group("/api/v1")
	
	{
		 public.POST("/auth/register", authHandler.Register)
		 public.POST("/auth/login", authHandler.Login)
		// public.GET("/health", healthHandler)
	}

	protected := router.Group("/api/v1")
	protected.Use(middleware.AuthMiddleware())
	
	
	{
		protected.POST("/order", orderHandler.PlaceOrder)
	}

	log.Println("API Gateway is running on port :8080")
	router.Run(":8080")
}