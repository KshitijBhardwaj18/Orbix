package handlers

import (
	"log"
	"time"

	"github.com/KshitijBhardwaj18/Orbix/services/api-gateway/types"
	"github.com/KshitijBhardwaj18/Orbix/shared/broker"
	"github.com/KshitijBhardwaj18/Orbix/shared/messages"
	"github.com/KshitijBhardwaj18/Orbix/shared/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type OrderHandler struct {
	broker *broker.Broker
}

func NewOrderHandler(brokerClient *broker.Broker) *OrderHandler {
	return &OrderHandler{broker: brokerClient}
}

func (h *OrderHandler) PlaceOrder(c *gin.Context) {
	var req struct {
		MarketID string `json:"market-id" binding:"required"`
		Side     string `json:"side" binding:"required,oneof=BUY SELL"`
		Type     string `json:"type" binding:"required,oneof=MARKET LIMIT"`
		Quantity string `json:"quantity" binding:"required"`
		Price    string `json:"price"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userIDstr := c.GetString("user_id")
	userID, err := uuid.Parse(userIDstr)

	if err != nil {
		c.JSON(405, gin.H{"error": "Invalid UserID formatt"})
	}

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
		UserID:   userID,
		MarketID: req.MarketID,
		Side:     models.OrderSide(req.Side),
		Type:     models.OrderType(req.Type),
		Quantity: quantity,
		Price:    price,
	}

	response, err := h.broker.CreateOrder(orderReq)

	orderResponse := types.OrderResponse{
        ID:                response.ID.String(),
        MarketID:          response.MarketID,
        Side:              string(response.Side),
        Type:              string(response.Type),
        Quantity:          response.Quantity.String(),
        FilledQuantity:    response.FilledQuantity.String(),
        RemainingQuantity: response.RemainingQuantity.String(),
        Status:            string(response.Status),
        CreatedAt:         response.CreatedAt.Format(time.RFC3339),
    }
    



	if err != nil {
		log.Printf("error: %v", err)
		c.JSON(500, gin.H{"error": "Failed to process order"})
		return
	}

	c.JSON(201, orderResponse)
}

func (h *OrderHandler) LogOrderbooks(c *gin.Context) {
	response, err := h.broker.LogOrderbooks()

	if err != nil {
		log.Printf("error logging orderbooks: %v", err)
		c.JSON(500, gin.H{"error": "Failed to retrieve orderbooks"})
		return
	}

	c.JSON(200, response)
}
