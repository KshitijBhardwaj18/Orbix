package models

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"time"
)

type Balance struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:idx_user_asset"`
	Asset string `gorm:"type:varchar(10);not null;uniqueIndex:idx_user_asset"`
	Available decimal.Decimal `gorm:"type:decimal(20,8);not null;default:0"`
	Locked decimal.Decimal `gorm:"type:decimal(20,8);not null;default:0"`
	CreatedAt time.Time
	UpdatedAt time.Time
	

	User User `gorm:"foreignKey:UserID"`
}