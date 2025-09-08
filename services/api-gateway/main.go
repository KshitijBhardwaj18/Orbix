package main

import (
	"log"

	"github.com/KshitijBhardwaj18/Orbix/services/api-gateway/config"
	"github.com/KshitijBhardwaj18/Orbix/services/api-gateway/handlers"
	"github.com/KshitijBhardwaj18/Orbix/services/api-gateway/middleware"
	"github.com/KshitijBhardwaj18/Orbix/shared/broker"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"time"
)

func main() {
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	Broker := broker.NewRedisClient()

    authHandler := handlers.NewAuthHandler(db)
	orderHandler := handlers.NewOrderHandler(Broker)
	userHandler := handlers.NewUserHandler(db)
	

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // your React frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))

	public := router.Group("/api/v1")
	
	{
		 public.POST("/auth/register", authHandler.Register)
		 public.POST("/auth/login", authHandler.Login)
		 public.POST("/auth/logout", authHandler.Logout)
	}

	protected := router.Group("/api/v1")
	protected.Use(middleware.AuthMiddleware())
	
	
	{
		protected.POST("/order", orderHandler.PlaceOrder)
		protected.GET("/user/me", userHandler.GetUser)
		
	}

	log.Println("API Gateway is running on port :8080")
	router.Run(":8080")
}