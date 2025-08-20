package models

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
    ID           uuid.UUID      `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
    Email        string         `gorm:"type:varchar(255);uniqueIndex;not null"`
    Username     string         `gorm:"type:varchar(50);uniqueIndex;not null"`
    PasswordHash string         `gorm:"type:varchar(255);not null"`
    CreatedAt    time.Time
    UpdatedAt    time.Time
    DeletedAt    gorm.DeletedAt `gorm:"index"`
    
    
    Orders   []Order   `gorm:"foreignKey:UserID"`
    Balances []Balance `gorm:"foreignKey:UserID"`
}

