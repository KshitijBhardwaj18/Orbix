package orderbook

import (
	"fmt"
	"log"
	"time"

	"github.com/KshitijBhardwaj18/Orbix/shared/models"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"sort"
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
	if order.Side == "BUY" {
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
				IsBuyerMaker:  true,
				Quantity:      filledQuantity,
				QuoteQuantity: o.Asks[i].Price.Mul(filledQuantity),
				CreatedAt:     time.Now(),
			}

			order.Trades = append(order.Trades, trade)
			o.CurrentPrice = trade.Price
			o.LastTradeId = trade.ID.String()

		}

	}

	for i := len(o.Asks) - 1; i >= 0; i-- {
		if o.Asks[i].RemainingQuantity.Equal(decimal.Zero) {
			log.Print("Reached here in the equation")
			o.Asks = append(o.Asks[:i], o.Asks[i+1:]...)

		}
	}

}

func (o *OrderBook) matchAsk(order *models.Order) {
	log.Printf("Reached here")
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
				IsBuyerMaker:  false,
			}

			order.Trades = append(order.Trades, trade)
			o.CurrentPrice = trade.Price
			o.LastTradeId = trade.ID.String()
		}
	}

	for i := len(o.Bids) - 1; i >= 0; i-- {
		if o.Bids[i].RemainingQuantity.Equal(decimal.Zero) {
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

func (o *OrderBook) GetDepth(maxLevels int) *MarketDepth {
	// Aggregate orders by price level
	bidLevels := o.aggregateOrdersByPrice(o.Bids)
	askLevels := o.aggregateOrdersByPrice(o.Asks)

	sort.Slice(bidLevels, func(i, j int) bool {
		return bidLevels[i].Price.GreaterThan(bidLevels[j].Price)
	})
	sort.Slice(askLevels, func(i, j int) bool {
		return askLevels[i].Price.LessThan(askLevels[j].Price)
	})

	bids := o.calculateCumulativeTotals(bidLevels, maxLevels)
	asks := o.calculateCumulativeTotals(askLevels, maxLevels)

	return &MarketDepth{
		Symbol:    o.GetTicker(),
		Bids:      bids,
		Asks:      asks,
		Timestamp: time.Now(),
	}
}

func (o *OrderBook) aggregateOrdersByPrice(orders []models.Order) []DepthLevel {
	priceMap := make(map[string]decimal.Decimal)

	for _, order := range orders {
		if order.Price == nil {
			continue
		}

		priceStr := order.Price.String()
		if existing, exists := priceMap[priceStr]; exists {
			priceMap[priceStr] = existing.Add(order.RemainingQuantity)
		} else {
			priceMap[priceStr] = order.RemainingQuantity
		}
	}

	levels := make([]DepthLevel, 0, len(priceMap))
	for priceStr, quantity := range priceMap {
		price, _ := decimal.NewFromString(priceStr)
		levels = append(levels, DepthLevel{
			Price:    price,
			Quantity: quantity,
		})
	}

	return levels
}

func (o *OrderBook) calculateCumulativeTotals(levels []DepthLevel, maxLevels int) []DepthLevel {
	if maxLevels > 0 && len(levels) > maxLevels {
		levels = levels[:maxLevels]
	}

	total := decimal.Zero

	for i := range levels {
		total = total.Add(levels[i].Quantity)
		levels[i].Total = total
	}

	return levels
}
