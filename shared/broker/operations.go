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

// func (r *RedisClient) PublishToClient(message interface{}, clientId string) error {
// 	return r.rdb.Publish(r.ctx,clientId,message).Err()
// }

type QueueMessage struct {
	ClientID string `json:"clientId"`
	MessageType string `json:"messageType"`
	Data interface{} `json:"data"`
}

func (r *RedisClient) BRPop(queueName string) (*QueueMessage, error){
	result, err := r.rdb.BRPop(r.ctx,0,queueName).Result()

	if err != nil {
		return nil, err
	}

	messageData := result[1]

	var queueMsg QueueMessage
	err = json.Unmarshal([]byte(messageData), &queueMsg)
	if err != nil {
		return nil, err
	}

	return &queueMsg, nil
}

func (r *RedisClient) PublishToClient(clientID string, response interface{}) error {
    data, err := json.Marshal(response)
    if err != nil {
        return err
    }
    
    return r.rdb.Publish(r.ctx, clientID, data).Err()
}
