package broker

import (
	"encoding/json"

	"github.com/KshitijBhardwaj18/Orbix/shared/messages"
	"github.com/google/uuid"
	"time"
	"errors"
)

func(r *RedisClient) CreateOrder(order *messages.OrderRequest) (*messages.OrderResponse, error) {
	clientId := uuid.New().String()

	request := map[string]interface{}{
		"clientId": clientId,
		"messageType": "CREATE_ORDER",
		"data": order,
	}

	pubsub := r.rdb.Subscribe(r.ctx,clientId)
	defer pubsub.Close()

	requestData, _ := json.Marshal(request)
	err := r.rdb.LPush(r.ctx,"engine_requests",requestData).Err()

	if err != nil {
		return nil, err
	}

	select {
	case msg := <-pubsub.Channel():
		var response messages.OrderResponse
		err := json.Unmarshal([]byte(msg.Payload), &response)
		return &response, err
	
	case <-time.After(5 * time.Second):
		return nil, errors.New("engine timeout")
	}
}