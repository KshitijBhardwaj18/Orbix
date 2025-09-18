package main

import (
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/KshitijBhardwaj18/Orbix/services/db/config"
	"github.com/KshitijBhardwaj18/Orbix/shared/broker"
	"github.com/KshitijBhardwaj18/Orbix/shared/models"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type EventProcessor struct {
	db     *gorm.DB
	broker *broker.Broker
}

func main() {
	// Initialize DB and broker
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	brokerInstance := broker.NewRedisClient()
	processor := &EventProcessor{db: db, broker: brokerInstance}

	log.Println("🚀 Starting Event Processor...")
	processor.Start()
}

func (ep *EventProcessor) Start() {
	// Subscribe to critical database events
	go ep.processOrderEvents()
	go ep.processTradeEvents()
	go ep.processTickerEvents() // 🎯 NEW: Process ticker events

	select {} // Keep running
}

func (ep *EventProcessor) processOrderEvents() {
	// Subscribe to order events
	pubsub := ep.broker.SubscribeToPattern("db@order*")
	defer pubsub.Close()

	for msg := range pubsub.Channel() {
		ep.handleOrderEvent(msg.Channel, msg.Payload)
	}
}

func (ep *EventProcessor) processTradeEvents() {
	pubsub := ep.broker.SubscribeToChannel("db@trade")
	defer pubsub.Close()

	for msg := range pubsub.Channel() {
		ep.handleTradeEvent(msg.Payload)
	}
}

// 🎯 NEW: Process ticker events
func (ep *EventProcessor) processTickerEvents() {
	pubsub := ep.broker.SubscribeToChannel("db@ticker")
	defer pubsub.Close()

	for msg := range pubsub.Channel() {
		ep.handleTickerEvent(msg.Payload)
	}
}

func (ep *EventProcessor) handleOrderEvent(channel, payload string) {
	var eventData map[string]interface{}
	if err := json.Unmarshal([]byte(payload), &eventData); err != nil {
		log.Printf("❌ Failed to parse order event: %v", err)
		return
	}

	orderData, _ := json.Marshal(eventData["order"])
	var order models.Order
	if err := json.Unmarshal(orderData, &order); err != nil {
		log.Printf("❌ Failed to parse order: %v", err)
		return
	}

	if strings.Contains(channel, "orderplaced") {
		// 🟢 INSERT new order
		if err := ep.db.Create(&order).Error; err != nil {
			log.Printf("❌ Failed to insert order: %v", err)
		} else {
			log.Printf("✅ Inserted order %s", order.ID.String())
		}
	} else {
		// 🟡 UPDATE existing order
		if err := ep.db.Save(&order).Error; err != nil {
			log.Printf("❌ Failed to update order: %v", err)
		} else {
			log.Printf("✅ Updated order %s (Status: %s)", order.ID.String(), order.Status)
		}
	}
}

func (ep *EventProcessor) handleTradeEvent(payload string) {
	var eventData map[string]interface{}
	if err := json.Unmarshal([]byte(payload), &eventData); err != nil {
		log.Printf("❌ Failed to parse trade event: %v", err)
		return
	}

	tradeData, _ := json.Marshal(eventData["trade"])
	var trade models.Trade
	if err := json.Unmarshal(tradeData, &trade); err != nil {
		log.Printf("❌ Failed to parse trade: %v", err)
		return
	}

	// 🟢 INSERT trade immediately
	if err := ep.db.Create(&trade).Error; err != nil {
		log.Printf("❌ Failed to insert trade: %v", err)
	} else {
		log.Printf("✅ Inserted trade %s (Price: %s, Quantity: %s)",
			trade.ID.String(), trade.Price.String(), trade.Quantity.String())
	}
}

// 🎯 NEW: Handle ticker events
func (ep *EventProcessor) handleTickerEvent(payload string) {
	var eventData map[string]interface{}
	if err := json.Unmarshal([]byte(payload), &eventData); err != nil {
		log.Printf("❌ Failed to parse ticker event: %v", err)
		return
	}

	market, _ := eventData["market"].(string)
	tickerStats, _ := eventData["ticker_stats"].(map[string]interface{})

	if market == "" || tickerStats == nil {
		log.Printf("❌ Invalid ticker event data")
		return
	}

	// Update market statistics in database
	ep.updateMarketStats(market, tickerStats)
}

func (ep *EventProcessor) updateMarketStats(marketID string, stats map[string]interface{}) {
	// Convert market ID format (BTC/USD -> BTCUSD)
	dbMarketID := strings.Replace(marketID, "/", "", 1)

	// Parse decimal values from strings
	currentPrice, _ := decimal.NewFromString(getString(stats, "current_price"))
	bestBid, _ := decimal.NewFromString(getString(stats, "best_bid"))
	bestAsk, _ := decimal.NewFromString(getString(stats, "best_ask"))
	spread, _ := decimal.NewFromString(getString(stats, "spread"))
	spreadPercent, _ := decimal.NewFromString(getString(stats, "spread_percent"))
	volume24h, _ := decimal.NewFromString(getString(stats, "volume_24h"))
	quoteVolume24h, _ := decimal.NewFromString(getString(stats, "quote_volume_24h"))
	high24h, _ := decimal.NewFromString(getString(stats, "high_24h"))
	low24h, _ := decimal.NewFromString(getString(stats, "low_24h"))
	priceChange24h, _ := decimal.NewFromString(getString(stats, "price_change_24h"))
	priceChangePercent24h, _ := decimal.NewFromString(getString(stats, "price_change_percent_24h"))

	// Update market in database
	updates := map[string]interface{}{
		"last_price":               currentPrice,
		"best_bid_price":           bestBid,
		"best_ask_price":           bestAsk,
		"spread":                   spread,
		"spread_percent":           spreadPercent,
		"volume24h":                volume24h,
		"quote_volume24h":          quoteVolume24h,
		"high_price24h":            high24h,
		"low_price24h":             low24h,
		"price_change24h":          priceChange24h,
		"price_change_percent24h":  priceChangePercent24h,
		"last_update_time":         time.Now(),
	}

	if err := ep.db.Model(&models.Market{}).Where("id = ?", dbMarketID).Updates(updates).Error; err != nil {
		log.Printf("❌ Failed to update market stats for %s: %v", marketID, err)
	} else {
		log.Printf("✅ Updated market stats for %s (Price: %s, Spread: %s)",
			marketID, currentPrice.String(), spread.String())
	}
}

// Helper function to safely get string from interface{}
func getString(m map[string]interface{}, key string) string {
	if val, ok := m[key].(string); ok {
		return val
	}
	return "0"
}
