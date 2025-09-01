package engine

import (
	"fmt"
	"time"

	"encoding/json"
	"github.com/KshitijBhardwaj18/Orbix/services/engine/orderbook"
	"github.com/KshitijBhardwaj18/Orbix/shared/broker"
	"github.com/KshitijBhardwaj18/Orbix/shared/messages"
	"github.com/KshitijBhardwaj18/Orbix/shared/models"
	"github.com/KshitijBhardwaj18/Orbix/shared/utils"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"log"
)

type UserBalances map[string]models.Balance
type BalanceCache map[uuid.UUID]UserBalances

type Engine struct {
	Orderbooks []*orderbook.OrderBook
	Balances   BalanceCache
	Broker     *broker.Broker
}

func NewEngine(broker *broker.Broker) *Engine {
	return &Engine{
		Orderbooks: []*orderbook.OrderBook{},
		Balances:   make(BalanceCache),
		Broker:     broker,
	}
}

func (e *Engine) Consume(message *messages.MessageFromAPI) {
	switch message.MessageType {
	case "CREATE_ORDER":
		dataBytes, _ := json.Marshal(message.Data)

		var orderReq messages.OrderRequest

		err := json.Unmarshal(dataBytes, &orderReq)

		if err != nil {
			log.Printf("Failed to parse order request: %v", err)
			return
		}

		order, err := e.CreateOrder(orderReq)

		if err != nil {
			fmt.Printf("error processing the order")
		}

		e.Broker.PublishToClient("ORDER_UPDATE", message.ClientId, order)
	}
}

func (e *Engine) CreateOrder(orderRequest messages.OrderRequest) (order *models.Order, err error) {
	orderbook, err := e.FindOrCreateOrderbook(orderRequest.MarketID)

	if err != nil {
		return nil, err
	}
	orderID := uuid.New()
	order = &models.Order{
		ID:                orderID,
		UserID:            orderRequest.UserID,
		MarketID:          orderRequest.MarketID,
		Side:              orderRequest.Side,
		Type:              orderRequest.Type,
		Quantity:          orderRequest.Quantity,
		Price:             orderRequest.Price,
		FilledQuantity:    decimal.Zero,
		RemainingQuantity: orderRequest.Quantity,
		Status:            models.PENDING,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	order = orderbook.AddOrder(order)

	e.LogOrderbooks()

	return order, nil
}

func (e *Engine) FindOrCreateOrderbook(marketID string) (orderBook *orderbook.OrderBook, err error) {
	for i := range e.Orderbooks {
		if e.Orderbooks[i].GetTicker() == marketID {
			return e.Orderbooks[i], nil
		}
	}

	baseAsset, QuoteAsset, _ := utils.ParseMarketId(marketID)

	orderbook := orderbook.NewOrderBook(baseAsset, QuoteAsset)

	e.Orderbooks = append(e.Orderbooks, orderbook)

	return orderbook, nil

}

func (e *Engine) LogOrderbooks() {
	fmt.Printf("Total orderbooks: %d\n", len(e.Orderbooks))
	for i, ob := range e.Orderbooks {
		fmt.Printf("Orderbook %d: %s (Bids: %d, Asks: %d)\n",
			i, ob.GetTicker(), len(ob.Bids), len(ob.Asks))
	}
}
