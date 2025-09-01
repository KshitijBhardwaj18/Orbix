package engine

import (
	"github.com/KshitijBhardwaj18/Orbix/services/engine/orderbook"
	"github.com/KshitijBhardwaj18/Orbix/shared/messages"
	"github.com/KshitijBhardwaj18/Orbix/shared/models"
	"github.com/KshitijBhardwaj18/Orbix/shared/utils"
	"github.com/google/uuid"
)

type UserBalances map[string]models.Balance
type BalanceCache map[uuid.UUID]UserBalances

type Engine struct {
	Orderbooks []orderbook.OrderBook
	Balances   BalanceCache
}

func (e *Engine) Consume(message *messages.MessageFromAPI, clientID string) {
	switch message.MessageType {
	case "CREATE_ORDER":
		orderData := message.Data.(*messages.OrderRequest)

		order, err := e.CreateOrder(*orderData)
	}
}

func (e *Engine) CreateOrder(orderRequest messages.OrderRequest) (order *models.Order, err error) {
	baseAsset, QuoteAsset, _ := utils.ParseMarketID(orderRequest.MarketID)

}
