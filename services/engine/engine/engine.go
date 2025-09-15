package engine

import (
	"fmt"
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

// Realistic market data for seeding
var MarketSeedingData = []MarketSeedData{
	{AvailableMarkets[0], 111227.70, 0.0005, 10}, // BTC/USD
	{AvailableMarkets[1], 4296.38, 0.0008, 10},   // ETH/USD
	{AvailableMarkets[2], 1.0001, 0.0001, 8},     // USDT/USD
	{AvailableMarkets[3], 207.27, 0.0012, 10},    // SOL/USD
	{AvailableMarkets[4], 0.23229, 0.0015, 8},    // DOGE/USD
	{AvailableMarkets[5], 22.427, 0.0010, 8},     // LINK/USD
	{AvailableMarkets[6], 3.397, 0.0015, 8},      // SUI/USD
	{AvailableMarkets[7], 0.00001255, 0.0020, 6}, // SHIB/USD
	{AvailableMarkets[8], 3.572, 0.0015, 6},      // RENDER/USD
	{AvailableMarkets[9], 0.29544, 0.0018, 6},    // SEI/USD
	{AvailableMarkets[10], 0.9155, 0.0012, 6},    // ONDO/USD
	{AvailableMarkets[11], 1.2728, 0.0015, 6},    // WLD/USD
	{AvailableMarkets[12], 0.03131, 0.0025, 6},   // PENGU/USD
	{AvailableMarkets[13], 0.00001001, 0.0020, 6}, // PEPE/USD
	{AvailableMarkets[14], 4.322, 0.0015, 6},     // APT/USD
	{AvailableMarkets[15], 0.2779, 0.0018, 6},    // POL/USD
	{AvailableMarkets[16], 9.40, 0.0012, 6},      // UNI/USD
	{AvailableMarkets[17], 0.765, 0.0015, 6},     // ENA/USD
	{AvailableMarkets[18], 299.59, 0.0010, 8},    // AAVE/USD
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
	log.Printf("🌱 Seeding markets with demo orders...")

	// Create a demo user ID for seeding orders
	demoUserID := uuid.New()

	for _, seedData := range MarketSeedingData {
		err := e.seedMarketOrders(seedData, demoUserID)
		if err != nil {
			log.Printf("Failed to seed market %s: %v", seedData.Market.Ticker, err)
			continue
		}
		log.Printf("✓ Seeded %s with %d bid/ask levels", seedData.Market.Ticker, seedData.Depth)
	}

	log.Printf("🎉 Market seeding completed! All markets ready for demo.")
	return nil
}

// seedMarketOrders creates realistic bid and ask orders for a specific market
func (e *Engine) seedMarketOrders(seedData MarketSeedData, userID uuid.UUID) error {
	orderbook, err := e.FindOrCreateOrderbook(seedData.Market.Ticker)
	if err != nil {
		return err
	}

	basePrice := decimal.NewFromFloat(seedData.BasePrice)
	spreadPercent := decimal.NewFromFloat(seedData.Spread)

	// Calculate bid and ask starting prices
	spread := basePrice.Mul(spreadPercent)
	bestBid := basePrice.Sub(spread.Div(decimal.NewFromInt(2)))
	bestAsk := basePrice.Add(spread.Div(decimal.NewFromInt(2)))

	// Create bid orders (buy orders below market price)
	for i := 0; i < seedData.Depth; i++ {
		// Price decreases as we go down the bid ladder
		priceReduction := decimal.NewFromFloat(0.0005 + float64(i)*0.0002) // 0.05% to 0.25%
		price := bestBid.Sub(basePrice.Mul(priceReduction))

		// Generate realistic quantity (varies by market)
		quantity := e.generateRealisticQuantity(seedData.Market.Ticker, price)

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
			CreatedAt:         time.Now().Add(-time.Duration(i) * time.Minute), // Stagger times
			UpdatedAt:         time.Now(),
		}

		orderbook.AddOrder(bidOrder)
	}

	// Create ask orders (sell orders above market price)
	for i := 0; i < seedData.Depth; i++ {
		// Price increases as we go up the ask ladder
		priceIncrease := decimal.NewFromFloat(0.0005 + float64(i)*0.0002) // 0.05% to 0.25%
		price := bestAsk.Add(basePrice.Mul(priceIncrease))

		// Generate realistic quantity
		quantity := e.generateRealisticQuantity(seedData.Market.Ticker, price)

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
			CreatedAt:         time.Now().Add(-time.Duration(i) * time.Minute),
			UpdatedAt:         time.Now(),
		}

		orderbook.AddOrder(askOrder)
	}

	return nil
}

// generateRealisticQuantity creates realistic order quantities based on the market
func (e *Engine) generateRealisticQuantity(ticker string, price decimal.Decimal) decimal.Decimal {
	// Seed random number generator based on ticker for consistent quantities
	hash := 0
	for _, char := range ticker {
		hash = hash*31 + int(char)
	}

	// Create different quantity ranges based on asset type
	switch {
	case strings.Contains(ticker, "BTC"):
		// Bitcoin: smaller quantities (0.01 to 2.5 BTC)
		base := 0.01 + float64(hash%100)/100.0*2.49
		return decimal.NewFromFloat(base)

	case strings.Contains(ticker, "ETH"):
		// Ethereum: medium quantities (0.1 to 25 ETH)
		base := 0.1 + float64(hash%150)/150.0*24.9
		return decimal.NewFromFloat(base)

	case strings.Contains(ticker, "USDT") || strings.Contains(ticker, "USD"):
		// Stablecoins: larger quantities (100 to 50000)
		base := 100 + float64(hash%200)*248.5
		return decimal.NewFromFloat(base)

	case strings.Contains(ticker, "SHIB") || strings.Contains(ticker, "PEPE"):
		// Meme coins: very large quantities
		base := 1000000 + float64(hash%300)*16666.67
		return decimal.NewFromFloat(base)

	case strings.Contains(ticker, "DOGE"):
		// Dogecoin: large quantities
		base := 1000 + float64(hash%250)*198
		return decimal.NewFromFloat(base)

	default:
		// Other altcoins: medium-large quantities (1 to 500)
		base := 1 + float64(hash%200)*2.495
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

	case "GET_MARKETS":
		markets := e.GetAllMarkets()
		e.Broker.PublishToClient("MARKETS", message.ClientId, markets)
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

	depth := orderbook.GetDepthResponse(10)

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