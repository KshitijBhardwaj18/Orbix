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
    
    // Market Data Fields - Essential for any trading platform
    LastPrice          decimal.Decimal `gorm:"type:decimal(20,8);default:0"`        // Current market price
    BestBidPrice       decimal.Decimal `gorm:"type:decimal(20,8);default:0"`        // Highest buy order price
    BestAskPrice       decimal.Decimal `gorm:"type:decimal(20,8);default:0"`        // Lowest sell order price
    
    // 24-hour Statistics
    HighPrice24h       decimal.Decimal `gorm:"type:decimal(20,8);default:0"`        // 24h high
    LowPrice24h        decimal.Decimal `gorm:"type:decimal(20,8);default:0"`         // 24h low
    Volume24h          decimal.Decimal `gorm:"type:decimal(30,8);default:0"`        // 24h base volume (e.g., BTC volume)
    QuoteVolume24h     decimal.Decimal `gorm:"type:decimal(30,8);default:0"`        // 24h quote volume (e.g., USD volume)
    PriceChange24h     decimal.Decimal `gorm:"type:decimal(20,8);default:0"`        // 24h price change (absolute)
    PriceChangePercent24h decimal.Decimal `gorm:"type:decimal(10,4);default:0"`      // 24h price change (percentage)
    
    // Real-time Market Stats
    Spread             decimal.Decimal `gorm:"type:decimal(20,8);default:0"`        // Ask - Bid
    SpreadPercent      decimal.Decimal `gorm:"type:decimal(10,4);default:0"`        // Spread as percentage
    TradeCount24h      int64           `gorm:"default:0"`                           // Number of trades in 24h
    
    // Market Status
    LastTradeTime      *time.Time      `gorm:"default:null"`                        // When last trade occurred
    LastUpdateTime     time.Time       `gorm:"default:CURRENT_TIMESTAMP"`           // When market data was last updated
    
    CreatedAt          time.Time
    UpdatedAt          time.Time
    
    // Relationships
    Orders []Order `gorm:"foreignKey:MarketID"`
    Trades []Trade `gorm:"foreignKey:MarketID"`
}

// MarketTicker represents the ticker data typically sent to clients
type MarketTicker struct {
    Symbol             string          `json:"symbol"`              // BTC/USD
    LastPrice          decimal.Decimal `json:"lastPrice"`           // Current price
    BestBid            decimal.Decimal `json:"bestBid"`             // Best bid price
    BestAsk            decimal.Decimal `json:"bestAsk"`             // Best ask price
    HighPrice24h       decimal.Decimal `json:"high24h"`             // 24h high
    LowPrice24h        decimal.Decimal `json:"low24h"`              // 24h low
    Volume24h          decimal.Decimal `json:"volume24h"`           // 24h volume
    QuoteVolume24h     decimal.Decimal `json:"quoteVolume24h"`      // 24h quote volume
    PriceChange24h     decimal.Decimal `json:"priceChange24h"`      // 24h change
    PriceChangePercent24h decimal.Decimal `json:"priceChangePercent24h"` // 24h change %
    TradeCount24h      int64           `json:"tradeCount24h"`       // 24h trade count
    Timestamp          time.Time       `json:"timestamp"`           // Data timestamp
}

// ToTicker converts Market to MarketTicker for API responses
func (m *Market) ToTicker() MarketTicker {
    return MarketTicker{
        Symbol:             m.ID,
        LastPrice:          m.LastPrice,
        BestBid:            m.BestBidPrice,
        BestAsk:            m.BestAskPrice,
        HighPrice24h:       m.HighPrice24h,
        LowPrice24h:        m.LowPrice24h,
        Volume24h:          m.Volume24h,
        QuoteVolume24h:     m.QuoteVolume24h,
        PriceChange24h:     m.PriceChange24h,
        PriceChangePercent24h: m.PriceChangePercent24h,
        TradeCount24h:      m.TradeCount24h,
        Timestamp:          m.LastUpdateTime,
    }
}