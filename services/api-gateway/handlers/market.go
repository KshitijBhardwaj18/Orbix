package handlers

import (
	"github.com/KshitijBhardwaj18/Orbix/shared/broker"
	"github.com/gin-gonic/gin"
)

type MarketHandler struct {
	broker *broker.Broker
}

func NewMarketHandler(broker *broker.Broker) *MarketHandler {
	return &MarketHandler{broker:broker}
}

func (h *MarketHandler) GetDepth(c *gin.Context) {
	market :=  c.Param("market")

	if market == "" {
		c.JSON(400, gin.H{"error": "Market parameter is required"})
		return
	}

	
}
