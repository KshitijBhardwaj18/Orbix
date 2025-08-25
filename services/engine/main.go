package main

import (
	"log"
	"github.com/KshitijBhardwaj18/Orbix/shared/broker"
)

func main() {
	
	Broker := broker.NewRedisClient();
	log.Println("enigne working")

	for{
		request,_ := Broker.BRPop("engine_requests")
		
		
		log.Printf("worked")
		res := "working"
		broker.NewRedisClient().PublishToClient(request.ClientID,res)

	}

	
}