package models

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"time"
)

type Order struct {
	ID                uuid.UUID        `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserID            uuid.UUID        `gorm:"type:uuid;not null;index"`
	MarketID          string           `gorm:"type:varchar(20);not null;index"`
	Side              OrderSide        `gorm:"type:varchar(4);not null"`
	Type              OrderType        `gorm:"type:varchar(10);not null"`
	Quantity          decimal.Decimal  `gorm:"type:decimal(20,8);not null"`
	Price             *decimal.Decimal `gorm:"type:decimal(20,8)"`
	FilledQuantity    decimal.Decimal  `gorm:"type:decimal(20,8);default:0"`
	RemainingQuantity decimal.Decimal  `gorm:"type:decimal(20,8);not null"`
	Status            OrderStatus      `gorm:"type:varchar(10);default:'PENDING'"`
	CreatedAt         time.Time
	UpdatedAt         time.Time

	// Relationships
	User   User    `gorm:"foreignKey:UserID"`
	Market Market  `gorm:"foreignKey:MarketID"`
	Trades []Trade `gorm:"foreignKey:BuyerOrderID;foreignKey:SellerOrderID"`
}

type OrderSide string

const (
	BUY  OrderSide = "BUY"
	SELL OrderSide = "SELL"
)

type OrderType string

const (
	MARKET OrderType = "MARKET"
	LIMIT  OrderType = "LIMIT"
)

type OrderStatus string

const (
	PENDING   OrderStatus = "PENDING"
	FILLED    OrderStatus = "FILLED"
	PARTIAL   OrderStatus = "PARTIAL"
	CANCELLED OrderStatus = "CANCELLED"
	REJECTED  OrderStatus = "REJECTED"
)
