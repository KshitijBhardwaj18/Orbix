package handlers

import (
	"log"
	"time"

	"github.com/KshitijBhardwaj18/Orbix/shared/broker"
	"github.com/KshitijBhardwaj18/Orbix/shared/messages"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

type OrderHandler struct {
	broker *broker.RedisClient
}

func NewOrderHandler(brokerClient *broker.RedisClient) *OrderHandler{
	return &OrderHandler{broker:brokerClient}
} 

func (h *OrderHandler) PlaceOrder(c *gin.Context){
	var req struct {
		Symbol     string `json:"symbol" binding:"required"`
		Side       string `json:"side" binding:"required,oneof=BUY SELL"`
		Type       string `json:"type" binding:"required,oneof=MARKET LIMIT"`
		Quantity   string `json:"quantity" binding:"required"`
		Price	   string `json:"price"`	
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetString("user_id")

	quantity, err := decimal.NewFromString(req.Quantity)

	if err != nil || quantity.LessThanOrEqual(decimal.Zero) {
		c.JSON(400, gin.H{"error": "Invalid quantity"})
		return 
	}

	var price *decimal.Decimal
    if req.Type == "LIMIT" {
        if req.Price == "" {
            c.JSON(400, gin.H{"error": "Price required for limit orders"})
            return
        }
        
        p, err := decimal.NewFromString(req.Price)
        if err != nil || p.LessThanOrEqual(decimal.Zero) {
            c.JSON(400, gin.H{"error": "Invalid price"})
            return
        }
        price = &p
    }

	orderReq := &messages.OrderRequest{
		UserID: userID,
		Symbol: req.Symbol,
		Side: req.Side,
		Type: req.Type,
		Quantity: quantity,
		Price: price,
		Timestamp: time.Now(),
	}

	response,err := h.broker.CreateOrder(orderReq); 
	if err != nil {
		log.Printf("error: %v", err)
		c.JSON(500, gin.H{"error": "Failed to process order"})
		return
	}

	c.JSON(201, response)
} 