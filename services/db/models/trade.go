package models 

import (
	"time"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Trade struct {
    ID            uuid.UUID       `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
    MarketID      string          `gorm:"type:varchar(20);not null;index"`
    BuyerOrderID  uuid.UUID       `gorm:"type:uuid;not null"`
    SellerOrderID uuid.UUID       `gorm:"type:uuid;not null"`
    BuyerID       uuid.UUID       `gorm:"type:uuid;not null"`
    SellerID      uuid.UUID       `gorm:"type:uuid;not null"`
    Price         decimal.Decimal `gorm:"type:decimal(20,8);not null"`
    Quantity      decimal.Decimal `gorm:"type:decimal(20,8);not null"`
    QuoteQuantity decimal.Decimal `gorm:"type:decimal(20,8);not null"`
	BuyerFee     decimal.Decimal `gorm:"type:decimal(20,8);not null"`
	SellerFee    decimal.Decimal `gorm:"type:decimal(20,8);not null"`
	IsBuyerMaker  bool            `gorm:"not null"`
	CreatedAt    time.Time       `gorm:"index"`


    Market      Market `gorm:"foreignKey:MarketID"`
    BuyerOrder  Order  `gorm:"foreignKey:BuyerOrderID"`
    SellerOrder Order  `gorm:"foreignKey:SellerOrderID"`
    Buyer       User   `gorm:"foreignKey:BuyerID"`
    Seller      User   `gorm:"foreignKey:SellerID"`
}