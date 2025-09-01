package main

import (
	"log"

	"github.com/KshitijBhardwaj18/Orbix/services/engine/engine"
	"github.com/KshitijBhardwaj18/Orbix/shared/broker"
)

func main() {
	Broker := broker.NewRedisClient()
	Engine := engine.NewEngine(Broker)

	for {
		message, err := Broker.BRPop("engine_requests")

		Engine.Consume(message)

		if err != nil {
			log.Printf("error is %v", err)
		}

	}

}
