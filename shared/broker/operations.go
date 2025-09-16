package broker

import (
	"encoding/json"

	"errors"
	"time"

	"github.com/KshitijBhardwaj18/Orbix/shared/messages"
	"github.com/KshitijBhardwaj18/Orbix/shared/models"
	"github.com/google/uuid"
)

func (r *Broker) CreateOrder(order *messages.OrderRequest) (*models.Order, error) {
	clientId := uuid.New().String()

	request := &messages.MessageFromAPI{
		ClientId:    clientId,
		MessageType: "CREATE_ORDER",
		Data:        order,
	}

	pubsub := r.rdb.Subscribe(r.ctx, clientId)
	defer pubsub.Close()

	requestData, _ := json.Marshal(request)
	err := r.rdb.LPush(r.ctx, "engine_requests", requestData).Err()

	if err != nil {
		return nil, err
	}

	select {
	case msg := <-pubsub.Channel():
		var response models.Order
		err := json.Unmarshal([]byte(msg.Payload), &response)
		return &response, err

	case <-time.After(5 * time.Second):
		return nil, errors.New("engine timeout")
	}
}

// OrderbookInfo represents individual orderbook data (same as in engine)
type OrderbookInfo struct {
	Ticker   string `json:"ticker"`
	BidCount int    `json:"bid_count"`
	AskCount int    `json:"ask_count"`
}

// OrderbooksResponse represents the complete orderbooks response (same as in engine)  
type OrderbooksResponse struct {
	TotalOrderbooks int             `json:"total_orderbooks"`
	Orderbooks      []OrderbookInfo `json:"orderbooks"`
}

func (r *Broker) LogOrderbooks() (*OrderbooksResponse, error){
	clientId := uuid.New().String()
	
	pubsub := r.rdb.Subscribe(r.ctx,clientId)
	defer pubsub.Close()

	request := &messages.MessageFromAPI{
		ClientId: clientId,
		MessageType: "LOG_ORDERBOOK",
	}

	requestData, _ := json.Marshal(request)

	err := r.rdb.LPush(r.ctx,"engine_requests",requestData).Err()

	if err != nil {
		return nil, err
	}

	select {
	case msg := <-pubsub.Channel():
		var response OrderbooksResponse
		err := json.Unmarshal([]byte(msg.Payload), &response)
		return &response, err

	case <-time.After(5 * time.Second):
		return nil, errors.New("engine timeout")
	}
}

// DepthResponse represents the complete orderbooks response (same as in engine)  
type DepthResponse struct {
	Market string     `json:"market"`
	Bids   [][2]string `json:"bids"`
	Asks   [][2]string `json:"asks"`
}

func (r *Broker) GetDepth(req *messages.GetDepthRequest) (*DepthResponse, error) {
	clientId := uuid.New().String()
	
	pubsub := r.rdb.Subscribe(r.ctx,clientId)
	defer pubsub.Close()

	request := &messages.MessageFromAPI{
		ClientId: clientId,
		MessageType: "GET_DEPTH",
		Data: req,
	}

	requestData, _ := json.Marshal(request)

	err := r.rdb.LPush(r.ctx,"engine_requests",requestData).Err()

	if err != nil {
		return nil, err
	}

	select {
	case msg := <-pubsub.Channel():
		var response DepthResponse
		err := json.Unmarshal([]byte(msg.Payload), &response)
		return &response, err

	case <-time.After(5 * time.Second):
		return nil, errors.New("engine timeout")
	}
}

func (r *Broker) GetOpenOrders(req *messages.GetOpenOrdersRequest) ([]models.Order, error) {
	clientId := uuid.New().String()
	
	pubsub := r.rdb.Subscribe(r.ctx, clientId)
	defer pubsub.Close()

	request := &messages.MessageFromAPI{
		ClientId:    clientId,
		MessageType: "GET_OPEN_ORDERS",
		Data:        req,
	}

	requestData, _ := json.Marshal(request)
	err := r.rdb.LPush(r.ctx, "engine_requests", requestData).Err()

	if err != nil {
		return nil, err
	}

	select {
	case msg := <-pubsub.Channel():
		var response []models.Order
		err := json.Unmarshal([]byte(msg.Payload), &response)
		return response, err

	case <-time.After(5 * time.Second):
		return nil, errors.New("engine timeout")
	}
}

func (r *Broker) BRPop(queueName string) (*messages.MessageFromAPI, error) {
	result, err := r.rdb.BRPop(r.ctx, 0, queueName).Result()

	if err != nil {
		return nil, err
	}

	messageData := result[1]

	var queueMsg messages.MessageFromAPI
	err = json.Unmarshal([]byte(messageData), &queueMsg)
	if err != nil {
		return nil, err
	}

	return &queueMsg, nil
}

func (r *Broker) PublishToClient(Type string, clientID string, response interface{}) error {
	data, err := json.Marshal(response)
	if err != nil {
		return err
	}

	return r.rdb.Publish(r.ctx, clientID, data).Err()
}