package engine

import (
	"fmt"
	"time"

	"encoding/json"
	"log"

	"github.com/KshitijBhardwaj18/Orbix/services/engine/orderbook"
	"github.com/KshitijBhardwaj18/Orbix/shared/broker"
	"github.com/KshitijBhardwaj18/Orbix/shared/messages"
	"github.com/KshitijBhardwaj18/Orbix/shared/models"
	"github.com/KshitijBhardwaj18/Orbix/shared/utils"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type UserBalances map[string]models.Balance
type BalanceCache map[uuid.UUID]UserBalances
type Market struct {
	Name   string `json:"name"`
	Ticker string `json:"ticker"`
}

var AvailableMarkets = []Market{
	{"Bitcoin", "BTC/USD"},
	{"Ethereum", "ETH/USD"},
	{"USDT", "USDT/USD"},
	{"Solana", "SOL/USD"},
	{"Dogecoin", "DOGE/USD"},
	{"Chainlink", "LINK/USD"},
	{"Sui", "SUI/USD"},
	{"Shiba Inu", "SHIB/USD"},
	{"Render", "RENDER/USD"},
	{"Sei", "SEI/USD"},
	{"Ondo", "ONDO/USD"},
	{"Worldcoin", "WLD/USD"},
	{"Pudgy Penguins", "PENGU/USD"},
	{"Pepe", "PEPE/USD"},
	{"Aptos", "APT/USD"},
	{"POL (ex-MATIC)", "POL/USD"},
	{"Uniswap", "UNI/USD"},
	{"Ethena", "ENA/USD"},
	{"Aave", "AAVE/USD"},
}


type Engine struct {
	Orderbooks []*orderbook.OrderBook
	Markets	   []Market
	Balances   BalanceCache
	Broker     *broker.Broker
}

func NewEngine(broker *broker.Broker) *Engine {
	engine := &Engine{
		Orderbooks: []*orderbook.OrderBook{},
		Markets:    AvailableMarkets,
		Balances:   make(BalanceCache),
		Broker:     broker,
	}

	err := engine.InitializeMarketOrderbooks()
	if err != nil {
		log.Printf("Warning: Failed to initialize some market orderbooks: %v", err)
	}
	
	return engine
}

func (e *Engine) InitializeMarketOrderbooks() error {
	log.Printf("Initializing orderbooks for %d predefined markets...", len(e.Markets))
	
	for _, market := range e.Markets {
		_, err := e.FindOrCreateOrderbook(market.Ticker)
		if err != nil {
			log.Printf("Failed to initialize orderbook for %s (%s): %v", market.Name, market.Ticker, err)
			return err
		}
		log.Printf("✓ Initialized orderbook for %s (%s)", market.Name, market.Ticker)
	}
	
	log.Printf("Successfully initialized %d orderbooks", len(e.Orderbooks))
	return nil
}

func (e *Engine) GetAllMarkets() []Market {
	return e.Markets
}

func (e *Engine) GetMarketByTicker(ticker string) *Market {
	for _, market := range e.Markets {
		if market.Ticker == ticker {
			return &market
		}
	}
	return nil
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

		

	case "LOG_ORDERBOOK":
		response := e.LogOrderbooks()
		e.Broker.PublishToClient("ORDERBOOK_LOG", message.ClientId, response)

	case "GET_DEPTH":
		dataBytes, _ := json.Marshal(message.Data)

		var getDepthReq messages.GetDepthRequest

		err := json.Unmarshal(dataBytes, &getDepthReq)

		if err != nil {
			log.Printf("Failed to parse getDepth request")
		}

		depth := e.GetDepth(getDepthReq.Market)

		e.Broker.PublishToClient("DEPTH",message.ClientId,depth)


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

	// Log orderbooks for debugging (you might want to remove this in production)
	_ = e.LogOrderbooks()

	return order, nil
}



func (e *Engine) GetDepth(Market string) *messages.DepthResponse {

	orderbook, err := e.FindOrCreateOrderbook(Market)

	if err != nil {
		log.Printf("Error occured while finding orderbook")
	}
    
	depth := orderbook.GetDepthResponse(10);

	return depth

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

type OrderbookInfo struct {
	Ticker   string `json:"ticker"`
	BidCount int    `json:"bid_count"`
	AskCount int    `json:"ask_count"`
}

// OrderbooksResponse represents the complete orderbooks response
type OrderbooksResponse struct {
	TotalOrderbooks int             `json:"total_orderbooks"`
	Orderbooks      []OrderbookInfo `json:"orderbooks"`
}

func (e *Engine) LogOrderbooks() *OrderbooksResponse {
	orderbooks := make([]OrderbookInfo, len(e.Orderbooks))
	
	for i, ob := range e.Orderbooks {
		orderbooks[i] = OrderbookInfo{
			Ticker:   ob.GetTicker(),
			BidCount: len(ob.Bids),
			AskCount: len(ob.Asks),
		}
	}
	
	return &OrderbooksResponse{
		TotalOrderbooks: len(e.Orderbooks),
		Orderbooks:      orderbooks,
	}
}

