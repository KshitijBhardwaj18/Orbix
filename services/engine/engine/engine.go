package engine

import (
	"fmt"
	"math/rand"
	"strings"
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

// MarketSeedData contains realistic pricing data for seeding
type MarketSeedData struct {
	Market    Market
	BasePrice float64 // Current market price
	Spread    float64 // Bid-ask spread percentage (e.g., 0.001 = 0.1%)
	Depth     int     // Number of levels on each side
}

// HIGH LIQUIDITY market data for realistic demo
var MarketSeedingData = []MarketSeedData{
	{AvailableMarkets[0], 111227.70, 0.0002, 50}, // BTC/USD - 50 levels each side
	{AvailableMarkets[1], 4296.38, 0.0003, 40},   // ETH/USD - 40 levels each side
	{AvailableMarkets[2], 1.0001, 0.0001, 30},    // USDT/USD - 30 levels each side
	{AvailableMarkets[3], 207.27, 0.0004, 35},    // SOL/USD - 35 levels each side
	{AvailableMarkets[4], 0.23229, 0.0005, 25},   // DOGE/USD - 25 levels each side
	{AvailableMarkets[5], 22.427, 0.0003, 30},    // LINK/USD - 30 levels each side
	{AvailableMarkets[6], 3.397, 0.0004, 25},     // SUI/USD - 25 levels each side
	{AvailableMarkets[7], 0.00001255, 0.0008, 20}, // SHIB/USD - 20 levels each side
	{AvailableMarkets[8], 3.572, 0.0005, 20},     // RENDER/USD - 20 levels each side
	{AvailableMarkets[9], 0.29544, 0.0006, 20},   // SEI/USD - 20 levels each side
	{AvailableMarkets[10], 0.9155, 0.0004, 20},   // ONDO/USD - 20 levels each side
	{AvailableMarkets[11], 1.2728, 0.0005, 20},   // WLD/USD - 20 levels each side
	{AvailableMarkets[12], 0.03131, 0.0008, 15},  // PENGU/USD - 15 levels each side
	{AvailableMarkets[13], 0.00001001, 0.0008, 15}, // PEPE/USD - 15 levels each side
	{AvailableMarkets[14], 4.322, 0.0005, 20},    // APT/USD - 20 levels each side
	{AvailableMarkets[15], 0.2779, 0.0006, 20},   // POL/USD - 20 levels each side
	{AvailableMarkets[16], 9.40, 0.0004, 25},     // UNI/USD - 25 levels each side
	{AvailableMarkets[17], 0.765, 0.0005, 20},    // ENA/USD - 20 levels each side
	{AvailableMarkets[18], 299.59, 0.0003, 35},   // AAVE/USD - 35 levels each side
}

type Engine struct {
	Orderbooks []*orderbook.OrderBook
	Markets    []Market
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

	// Seed markets with demo orders for frontend demo
	err = engine.SeedMarketsWithOrders()
	if err != nil {
		log.Printf("Warning: Failed to seed some markets: %v", err)
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

// SeedMarketsWithOrders creates realistic demo orders for all markets
func (e *Engine) SeedMarketsWithOrders() error {
	log.Printf("🌱 Seeding markets with HIGH LIQUIDITY demo orders...")

	// Create multiple demo user IDs for variety
	demoUsers := make([]uuid.UUID, 5)
	for i := range demoUsers {
		demoUsers[i] = uuid.New()
	}

	totalOrders := 0
	for _, seedData := range MarketSeedingData {
		orders := e.seedMarketOrders(seedData, demoUsers)
		totalOrders += orders
		log.Printf("✓ Seeded %s with %d levels (%d total orders)", 
			seedData.Market.Ticker, seedData.Depth, orders)
	}

	log.Printf("🎉 HIGH LIQUIDITY seeding completed! %d total orders across all markets.", totalOrders)
	return nil
}

// seedMarketOrders creates realistic bid and ask orders for a specific market
func (e *Engine) seedMarketOrders(seedData MarketSeedData, userIDs []uuid.UUID) int {
	orderbook, err := e.FindOrCreateOrderbook(seedData.Market.Ticker)
	if err != nil {
		log.Printf("Failed to seed market %s: %v", seedData.Market.Ticker, err)
		return 0
	}

	basePrice := decimal.NewFromFloat(seedData.BasePrice)
	spreadPercent := decimal.NewFromFloat(seedData.Spread)

	// Calculate bid and ask starting prices
	spread := basePrice.Mul(spreadPercent)
	bestBid := basePrice.Sub(spread.Div(decimal.NewFromInt(2)))
	bestAsk := basePrice.Add(spread.Div(decimal.NewFromInt(2)))

	orderCount := 0

	// Create bid orders (buy orders below market price) with HIGH LIQUIDITY
	for i := 0; i < seedData.Depth; i++ {
		// Much smaller price steps for tighter liquidity
		priceReduction := decimal.NewFromFloat(0.0001 + float64(i)*0.00005) // 0.01% to 0.26%
		price := bestBid.Sub(basePrice.Mul(priceReduction))

		// Create 2-4 orders at each price level for deep liquidity
		ordersAtLevel := 2 + (i % 3) // 2-4 orders per level
		for j := 0; j < ordersAtLevel; j++ {
			quantity := e.generateHighLiquidityQuantity(seedData.Market.Ticker, price, j)
			userID := userIDs[j%len(userIDs)] // Rotate through users

			bidOrder := &models.Order{
				ID:                uuid.New(),
				UserID:            userID,
				MarketID:          seedData.Market.Ticker,
				Side:              models.BUY,
				Type:              models.LIMIT,
				Quantity:          quantity,
				Price:             &price,
				FilledQuantity:    decimal.Zero,
				RemainingQuantity: quantity,
				Status:            models.PENDING,
				CreatedAt:         time.Now().Add(-time.Duration(i*10+j) * time.Second),
				UpdatedAt:         time.Now(),
			}

			orderbook.AddOrder(bidOrder)
			orderCount++
		}
	}

	// Create ask orders (sell orders above market price) with HIGH LIQUIDITY
	for i := 0; i < seedData.Depth; i++ {
		// Much smaller price steps for tighter liquidity
		priceIncrease := decimal.NewFromFloat(0.0001 + float64(i)*0.00005) // 0.01% to 0.26%
		price := bestAsk.Add(basePrice.Mul(priceIncrease))

		// Create 2-4 orders at each price level for deep liquidity
		ordersAtLevel := 2 + (i % 3) // 2-4 orders per level
		for j := 0; j < ordersAtLevel; j++ {
			quantity := e.generateHighLiquidityQuantity(seedData.Market.Ticker, price, j)
			userID := userIDs[j%len(userIDs)] // Rotate through users

			askOrder := &models.Order{
				ID:                uuid.New(),
				UserID:            userID,
				MarketID:          seedData.Market.Ticker,
				Side:              models.SELL,
				Type:              models.LIMIT,
				Quantity:          quantity,
				Price:             &price,
				FilledQuantity:    decimal.Zero,
				RemainingQuantity: quantity,
				Status:            models.PENDING,
				CreatedAt:         time.Now().Add(-time.Duration(i*10+j) * time.Second),
				UpdatedAt:         time.Now(),
			}

			orderbook.AddOrder(askOrder)
			orderCount++
		}
	}

	return orderCount
}

// generateHighLiquidityQuantity creates much larger quantities for high liquidity appearance
func (e *Engine) generateHighLiquidityQuantity(ticker string, price decimal.Decimal, variation int) decimal.Decimal {
	// Use variation for different order sizes at same price level
	rand.Seed(int64(len(ticker)*1000 + variation*100))
	
	// Create different quantity ranges based on asset type with MUCH higher liquidity
	switch {
	case strings.Contains(ticker, "BTC"):
		// Bitcoin: 0.1 to 15 BTC per order (much higher than before)
		base := 0.1 + rand.Float64()*14.9
		return decimal.NewFromFloat(base)

	case strings.Contains(ticker, "ETH"):
		// Ethereum: 1 to 100 ETH per order (much higher)
		base := 1 + rand.Float64()*99
		return decimal.NewFromFloat(base)

	case strings.Contains(ticker, "USDT") || strings.Contains(ticker, "USD"):
		// Stablecoins: 1000 to 500000 (very high liquidity)
		base := 1000 + rand.Float64()*499000
		return decimal.NewFromFloat(base)

	case strings.Contains(ticker, "SHIB") || strings.Contains(ticker, "PEPE"):
		// Meme coins: 10M to 1B tokens (massive liquidity)
		base := 10000000 + rand.Float64()*990000000
		return decimal.NewFromFloat(base)

	case strings.Contains(ticker, "DOGE"):
		// Dogecoin: 10K to 1M DOGE (high liquidity)
		base := 10000 + rand.Float64()*990000
		return decimal.NewFromFloat(base)

	case strings.Contains(ticker, "SOL"):
		// Solana: 10 to 1000 SOL (high liquidity)
		base := 10 + rand.Float64()*990
		return decimal.NewFromFloat(base)

	default:
		// Other altcoins: 10 to 5000 tokens (much higher)
		base := 10 + rand.Float64()*4990
		return decimal.NewFromFloat(base)
	}
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

		e.Broker.PublishToClient("DEPTH", message.ClientId, depth)

	case "GET_OPEN_ORDERS":
		dataBytes, _ := json.Marshal(message.Data)

		log.Printf("Reached here")

		var getOpenOrdersReq messages.GetOpenOrdersRequest

		err := json.Unmarshal(dataBytes, &getOpenOrdersReq)

		if err != nil {
			log.Printf("Failed to parse getOpenOrders request: %v", err)
			return
		}

		orders := e.GetOpenOrders(getOpenOrdersReq.UserID, getOpenOrdersReq.Market)

		e.Broker.PublishToClient("OPEN_ORDERS", message.ClientId, orders)

	case "GET_MARKETS":
		markets := e.GetAllMarkets()
		e.Broker.PublishToClient("MARKETS", message.ClientId, markets)

	
	case "CANCEL_ORDER":
		dataBytes, _ := json.Marshal(message.Data)

		var cancelOrderRequest messages.CancelOrderRequest

		err := json.Unmarshal(dataBytes, &cancelOrderRequest)

		if err != nil {
			log.Printf("Failed to parse cancel order request: %v", err)
			return 
		}

		cancelledOrder, success := e.CancelOrder(cancelOrderRequest)
		
		// Prepare response
		if success && cancelledOrder != nil {
			e.Broker.PublishToClient("ORDER_CANCELLED", message.ClientId, messages.CancelOrderResponse{
				Success: true,
				Message: "Order cancelled successfully",
				OrderId: cancelOrderRequest.OrderID,
			})
		} else {
			e.Broker.PublishToClient("ORDER_CANCELLED", message.ClientId, messages.CancelOrderResponse{
				Success: false,
				Message: "Order cancellation failed",
				OrderId: "",
			})
		}
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

func (e *Engine) CancelOrder(req messages.CancelOrderRequest) (*models.Order, bool) {
	// Search through all orderbooks to find and cancel the order
	for _, orderbook := range e.Orderbooks {
		// Try to remove the order from this orderbook
		cancelledOrder, found := orderbook.RemoveOrder(req.OrderID, req.UserID)
		if found {
			log.Printf("✅ Order %s cancelled successfully for user %s in market %s", 
				req.OrderID, req.UserID.String(), orderbook.GetTicker())
			return cancelledOrder, true
		}
	}

	// Order not found in any orderbook
	log.Printf("❌ Order %s not found for user %s in any market", 
		req.OrderID, req.UserID.String())
	return nil, false
}

func (e *Engine) GetDepth(Market string) *messages.DepthResponse {

	orderbook, err := e.FindOrCreateOrderbook(Market)

	if err != nil {
		log.Printf("Error occured while finding orderbook")
	}

	depth := orderbook.GetDepthResponse(50) // Increased from 10 to 50 levels

	return depth
}

func (e *Engine) GetOpenOrders(userID uuid.UUID, market string) []models.Order {

	var openOrders []models.Order

	market = strings.Replace(market, "_", "/", 1)
	


	// If market is specified, only get orders for that market
	if market != "" {
		orderbook, err := e.FindOrCreateOrderbook(market)
		if err != nil {
			log.Printf("Error finding orderbook for market %s: %v", market, err)
			return openOrders
		}
		orders := orderbook.GetOpenOrders(userID)
		return orders
	}

	// If no market specified, get orders from all markets
	for _, ob := range e.Orderbooks {
		orders := ob.GetOpenOrders(userID)
		openOrders = append(openOrders, orders...)
	}

	return openOrders
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