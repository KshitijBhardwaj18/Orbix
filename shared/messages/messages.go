package messages

import (
	"time"

	"github.com/KshitijBhardwaj18/Orbix/shared/models"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type OrderRequest struct {
	UserID   uuid.UUID        `gorm:"type:uuid;not null;index"`
	MarketID string           `gorm:"type:varchar(20);not null;index"`
	Side     models.OrderSide `gorm:"type:varchar(4);not null"`
	Quantity decimal.Decimal  `gorm:"type:decimal(20,8);not null"`
	Price    *decimal.Decimal `gorm:"type:decimal(20,8)"`
	Type     models.OrderType `gorm:"type:varchar(5);not null"`
}

type CancelOrderRequest struct {
	UserID  string `json:"user_id"`
	OrderID string `json:"order_id"`
}

type OrderResponse struct {
	OrderID           string          `json:"order_id"`
	Status            string          `json:"status"`
	Message           string          `json:"message"`
	FilledQuantity    decimal.Decimal `json:"filled_quantity"`
	RemainingQuantity decimal.Decimal `json:"remaining_quantity"`
	Trades            []TradeEvent    `json:"trades"`
	Timestamp         time.Time       `json:"timestamp"`
}

type TradeEvent struct {
	TradeID     string `json:"trade_id"`
	Symbol      string `json:"symbol"`
	BuyOrderID  string `json:"buy_order_id"`
	SellOrderID string `json:"sell_order_id"`
	Price       string `json:"price"`
	Quantity    string `json:"quantity"`
	BuyerID     string `json:"buyer_id"`
	SellerID    string `json:"seller_id"`
	Timestamp   string `json:"timestamp"`
}

type OrderBookUpdate struct {
	Symbol string           `json:"symbol"`
	Bids   []OrderBookLevel `json:"bids"`
	Asks   []OrderBookLevel `json:"asks"`
}

type OrderBookLevel struct {
	Price    *decimal.Decimal `json:"price"`
	Quantity decimal.Decimal  `json:"quantity"`
}

type DepthLevel struct {
	Price    decimal.Decimal `json:"price"`
	Quantity decimal.Decimal `json:"quantity"`
	Total    decimal.Decimal `json:"total"`
}

type MarketDepth struct {
	Symbol    string       `json:"symbol"`
	Bids      []DepthLevel `json:"bids"`
	Asks      []DepthLevel `json:"asks"`
	Timestamp time.Time    `json:"timestamp"`
}

type MessageFromAPI struct {
	ClientId    string      `json:"clientId"`
	MessageType string      `json:"messageType"`
	Data        interface{} `json:"data"`
}
