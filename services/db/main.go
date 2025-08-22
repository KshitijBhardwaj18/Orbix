package main

import (
	"log"
	"github.com/KshitijBhardwaj18/Orbix/services/db/config"
	"github.com/KshitijBhardwaj18/Orbix/shared/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.ConnectDB()

	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)

	err = db.AutoMigrate(&models.User{}, &models.Market{}, &models.Order{}, &models.Trade{}, &models.Balance{})

	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database migrated successfully")
	log.Println("DB Service is running...")

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "healthy",
			"service": "db-service",
		})
	})

	log.Println("DB Service is running on port 8083")
	router.Run(":8083")
}
