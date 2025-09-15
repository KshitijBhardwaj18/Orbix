package handlers

import (
	"log"
	"strings"

	"github.com/KshitijBhardwaj18/Orbix/shared/broker"
	"github.com/KshitijBhardwaj18/Orbix/shared/messages"
	"github.com/gin-gonic/gin"
)

type MarketHandler struct {
	broker *broker.Broker
}

func NewMarketHandler(broker *broker.Broker) *MarketHandler {
	return &MarketHandler{broker:broker}
}

func (h *MarketHandler) GetDepth(c *gin.Context) {
	marketParam := c.Param("market")
	
	if marketParam == "" {
		c.JSON(400, gin.H{"error": "Market parameter is required"})
		return
	}

	// Convert BTC_USD format to BTC/USD format
	market := strings.Replace(marketParam, "_", "/", 1)

	req := &messages.GetDepthRequest{
		Market: market,
	}

	response, err := h.broker.GetDepth(req)

	if err != nil {
		log.Printf("Error in api receiving depth from engine: %v", err)
		c.JSON(500, gin.H{"error": "Failed to get market depth"})
		return
	}

	c.JSON(200, response)
}