package models

import (
    "time"
    "github.com/shopspring/decimal"
)

type Market struct {
    ID                 string          `gorm:"type:varchar(20);primaryKey"` // BTCUSDT
    BaseAsset          string          `gorm:"type:varchar(10);not null"`   // BTC
    QuoteAsset         string          `gorm:"type:varchar(10);not null"`   // USDT
    MinQuantity        decimal.Decimal `gorm:"type:decimal(20,8);not null;default:0.00000001"`
    MinPrice           decimal.Decimal `gorm:"type:decimal(20,8);not null;default:0.00000001"`
    PricePrecision     int             `gorm:"not null;default:8"`
    QuantityPrecision  int             `gorm:"not null;default:8"`
    IsActive           bool            `gorm:"not null;default:true"`
    CreatedAt          time.Time
    
    // Relationships
    Orders []Order `gorm:"foreignKey:MarketID"`
    Trades []Trade `gorm:"foreignKey:MarketID"`
}