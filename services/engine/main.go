package main

import (
	"log"

	"github.com/KshitijBhardwaj18/Orbix/shared/broker"
	"github.com/KshitijBhardwaj18/Orbix/shared/messages"
	"github.com/shopspring/decimal"
)

func main() {

	Broker := broker.NewRedisClient()
	log.Println("enigne working")

	for {
		request, err := Broker.BRPop("engine_requests")

		if err != nil {
			{
				log.Printf("BRPop error: %v", err)
				continue
			}
		}

		if request == nil {
			log.Printf("Recieved nil request")
			continue
		}

		log.Printf("worked")

		tradeEvent := messages.TradeEvent{
			TradeID:     "1234",
			Symbol:      "BTC/USD",
			BuyOrderID:  "1234",
			SellOrderID: "1234",
			Price:       "1234",
			Quantity:    "1234",
			BuyerID:     "1234",
			SellerID:    "12345",
			Timestamp:   "1234",
		}

		tradeEvent2 := messages.TradeEvent{
			TradeID:     "12345",
			Symbol:      "ETH/USD",
			BuyOrderID:  "1234",
			SellOrderID: "1234",
			Price:       "1234",
			Quantity:    "1234",
			BuyerID:     "1234",
			SellerID:    "12345",
			Timestamp:   "1234",
		}

		filledQty, _ := decimal.NewFromString("12344")
		remainingQty, _ := decimal.NewFromString("24234")

		res := messages.OrderResponse{
			OrderID:           "1234",
			Status:            "OK",
			Message:           "DEMO",
			FilledQuantity:    filledQty,
			RemainingQuantity: remainingQty,
			Trades:            []messages.TradeEvent{tradeEvent,tradeEvent2},
		}

		err = Broker.PublishToClient(request.ClientID, res)

		if err != nil {
			log.Printf("error is %v", err)
		}

	}

}
