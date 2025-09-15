package handlers

import (
	"log"

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
	market :=  c.Param("market")

	req := &messages.GetDepthRequest{
		Market: market,
	}

	if market == "" {
		c.JSON(400, gin.H{"error": "Market parameter is required"})
		return
	}

	response, err := h.broker.GetDepth(req)

	if err != nil {
		log.Printf("Error in api recivening depth from  engine")
	}

	c.JSON(200, response)
}
