package orderbook

import (
	"github.com/KshitijBhardwaj18/Orbix/shared/models"
	"github.com/shopspring/decimal"
	"fmt"
)

type OrderBook struct {
	BaseAsset    string
	QuoteAsset   string
	Bids         []models.Order
	Asks         []models.Order
	LastTradeId  string
	CurrentPrice decimal.Decimal
}

func NewOrderBook(BaseAsset, QuoteAsset string) *OrderBook {
	orderbook := OrderBook{BaseAsset: BaseAsset, 
						   QuoteAsset: QuoteAsset,
						   Bids: []models.Order{},
						   Asks: []models.Order{},
						   LastTradeId: "nil",
						   CurrentPrice: decimal.Zero,
						}
	return &orderbook
}

func (o *OrderBook) GetTicker() string{
	return fmt.Sprintf("%s/%s",o.BaseAsset,o.QuoteAsset)
}


