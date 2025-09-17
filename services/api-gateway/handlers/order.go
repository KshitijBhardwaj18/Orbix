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

	if response.Price != nil {
		orderResponse.Price = response.Price.String()
	}

	if err != nil {
		log.Printf("error: %v", err)
		c.JSON(500, gin.H{"error": "Failed to process order"})
		return
	}

	c.JSON(201, orderResponse)
}

func (h *OrderHandler) DeleteOrder(c *gin.Context) {
	log.Printf("🚀 DeleteOrder function started")
	
	var req struct {
		OrderId string `json:"order_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("❌ Failed to bind JSON: %v", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	log.Printf("📋 Request received - OrderId: %s", req.OrderId)

	userIdStr := c.GetString("user_id")
	log.Printf("🔍 User ID from context: '%s'", userIdStr)

	if userIdStr == "" {
		log.Printf("❌ No user_id found in context")
		c.JSON(401, gin.H{"error": "User not authenticated"})
		return
	}

	userID, err := uuid.Parse(userIdStr)
	if err != nil {
		log.Printf("❌ Failed to parse user ID '%s': %v", userIdStr, err)
		c.JSON(401, gin.H{"Authentication Error": "Error authenticating the request"})
		return
	}

	log.Printf("👤 Parsed user ID: %s", userID.String())

	var cancelReq messages.CancelOrderRequest = messages.CancelOrderRequest{
		UserID:  userID,
		OrderID: req.OrderId,
	}

	log.Printf("📤 Sending cancel request to broker - OrderID: %s, UserID: %s", 
		cancelReq.OrderID, cancelReq.UserID.String())

	response, err := h.broker.CancelOrderRequest(&cancelReq)

	if err != nil {
		log.Printf("❌ Broker error: %v", err)
		c.JSON(500, gin.H{"error": "Failed to cancel order", "details": err.Error()})
		return
	}

	log.Printf("📥 Broker response - Success: %t, Message: '%s', OrderId: '%s'", 
		response.Success, response.Message, response.OrderId)

	if response.Success {
		log.Printf("✅ Order %s cancelled successfully for user %s", 
			response.OrderId, userID.String())
		c.JSON(200, response)
	} else {
		log.Printf("⚠️ Order cancellation failed - %s", response.Message)
		c.JSON(400, response)
	}
}

func (h *OrderHandler) GetOpenOrders(c *gin.Context) {
	userIDstr := c.GetString("user_id")
	userID, err := uuid.Parse(userIDstr)

	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid user ID"})
		return
	}

	
	market := c.Query("market")

	req := &messages.GetOpenOrdersRequest{
		UserID: userID,
		Market: market, 
	}

	response, err := h.broker.GetOpenOrders(req)

	if err != nil {
		log.Printf("error getting open orders: %v", err)
		c.JSON(500, gin.H{"error": "Failed to retrieve open orders"})
		return
	}

	// Convert to API response format
	orderResponses := make([]types.OrderResponse, len(response))
	for i, order := range response {
		orderResponses[i] = types.OrderResponse{
			ID:                order.ID.String(),
			MarketID:          order.MarketID,
			Side:              string(order.Side),
			Type:              string(order.Type),
			Quantity:          order.Quantity.String(),
			FilledQuantity:    order.FilledQuantity.String(),
			RemainingQuantity: order.RemainingQuantity.String(),
			Status:            string(order.Status),
			CreatedAt:         order.CreatedAt.Format(time.RFC3339),
		}
		if order.Price != nil {
			orderResponses[i].Price = order.Price.String()
		}
	}

	c.JSON(200, gin.H{
		"orders": orderResponses,
		"count":  len(orderResponses),
	})
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

