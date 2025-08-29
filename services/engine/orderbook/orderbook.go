package orderbook

import (
	"fmt"
	"time"

	"github.com/KshitijBhardwaj18/Orbix/shared/models"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
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
		QuoteAsset:   QuoteAsset,
		Bids:         []models.Order{},
		Asks:         []models.Order{},
		LastTradeId:  "nil",
		CurrentPrice: decimal.Zero,
	}
	return &orderbook
}

func (o *OrderBook) GetTicker() string {
	return fmt.Sprintf("%s/%s", o.BaseAsset, o.QuoteAsset)
}

func (o *OrderBook) AddOrder(order *models.Order) *models.Order {
	if order.Side == "buy" {
		o.matchBid(order)
		if order.FilledQuantity == order.Quantity {
			return order
		}
		o.Bids = append(o.Bids, *order)
		return order
	} else {
		o.matchAsk(order)
		if order.FilledQuantity == order.Quantity {
			return order
		}
		o.Asks = append(o.Asks, *order)
		return order
	}
}

func (o *OrderBook) matchBid(order *models.Order) {
	for i := 0; i < len(o.Asks); i++ {
		if o.Asks[i].Price.LessThanOrEqual(*order.Price) && !order.FilledQuantity.Equal(order.Quantity) {
			filledQuantity := decimal.Min(o.Asks[i].RemainingQuantity, order.RemainingQuantity)

			order.FilledQuantity = order.FilledQuantity.Add(filledQuantity)
			order.RemainingQuantity = order.Quantity.Sub(order.FilledQuantity)

			o.Asks[i].FilledQuantity = o.Asks[i].FilledQuantity.Add(filledQuantity)
			o.Asks[i].RemainingQuantity = o.Asks[i].Quantity.Sub(o.Asks[i].FilledQuantity)

			trade := models.Trade{
				ID:            uuid.New(),
				MarketID:      order.MarketID,
				BuyerID:       order.UserID,
				SellerID:      o.Asks[i].UserID,
				BuyerOrderID:  order.ID,
				SellerOrderID: o.Asks[i].ID,
				Price:         *o.Asks[i].Price,
				IsBuyerMaker: true,
				Quantity:      filledQuantity,
				QuoteQuantity: o.Asks[i].Price.Mul(filledQuantity),
				CreatedAt: time.Now(),
			}

			order.Trades = append(order.Trades, trade)
			o.CurrentPrice = *o.Asks[i].Price
			o.LastTradeId = trade.ID.String()

		}

	}

	for i := 0; i < len(o.Asks); i++ {
		if o.Asks[i].FilledQuantity.Equal(o.Asks[i].Quantity) {
			o.Asks = append(o.Asks[:i], o.Asks[i+1:]...)
			i--
		}
	}

}

func (o *OrderBook) matchAsk(order *models.Order) {
	for i := 0; i < len(o.Bids); i++ {
		if o.Bids[i].Price.GreaterThanOrEqual(*order.Price) && !order.FilledQuantity.Equal(order.Quantity) {
			filledQuantity := decimal.Min(o.Bids[i].RemainingQuantity, order.RemainingQuantity)

			order.FilledQuantity = order.FilledQuantity.Add(filledQuantity)
			order.RemainingQuantity = order.Quantity.Sub(order.FilledQuantity)

			o.Bids[i].FilledQuantity = o.Bids[i].FilledQuantity.Add(filledQuantity)
			o.Bids[i].RemainingQuantity = o.Bids[i].Quantity.Sub(o.Bids[i].FilledQuantity)

			trade := models.Trade{
				ID:            uuid.New(),
				MarketID:      order.MarketID,
				BuyerID:       o.Bids[i].UserID,
				SellerID:      order.UserID,     
				BuyerOrderID:  o.Bids[i].ID,    
				SellerOrderID: order.ID,         
				Price:         *o.Bids[i].Price,     
				Quantity:      filledQuantity,
				QuoteQuantity: order.Price.Mul(filledQuantity),
				IsBuyerMaker: false,
			}

			order.Trades = append(order.Trades, trade)
			o.CurrentPrice = *o.Asks[i].Price
			o.LastTradeId = trade.ID.String()
		}
	}

	
	for i := len(o.Bids) - 1; i >= 0; i-- {
		if o.Bids[i].FilledQuantity.Equal(o.Bids[i].Quantity) {
			o.Bids = append(o.Bids[:i], o.Bids[i+1:]...)
		}
	}
}

func (o *OrderBook) GetOpenOrders(userID uuid.UUID) []models.Order {

    openOrders := make([]models.Order, 0, len(o.Asks)+len(o.Bids))
    
    
    for _, ask := range o.Asks {
        if ask.UserID == userID {
            openOrders = append(openOrders, ask)
        }
    }
    
    for _, bid := range o.Bids {
        if bid.UserID == userID {
            openOrders = append(openOrders, bid)
        }
    }
    
    return openOrders
}

func (o *OrderBook ) GetDepth() {
	
}
