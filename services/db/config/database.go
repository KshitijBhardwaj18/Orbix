package config

import (
	"fmt"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Host     string
	Port	 string
	User	 string
	Password string
	DBName	 string
}

func GetDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Host: getEnv("DB_HOST", "localhost"),
		Port: getEnv("DB_PORT", "5432"),
		User: getEnv("DB_USER", "admin"),
		Password: getEnv("DB_PASSWORD", "admin"),
		DBName: getEnv("DB_NAME", "obrbix"),
	}
}

func (config *DatabaseConfig) getDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.DBName)
}

func ConnectDB() (*gorm.DB,error) {
	config := GetDatabaseConfig()
	dsn := config.getDSN()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return db, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value 
	}

	return defaultValue
}

