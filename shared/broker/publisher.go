package broker

import (
	"encoding/json"

	"github.com/KshitijBhardwaj18/Orbix/shared/messages"
)

func (r *RedisClient) PublishOrder(order *messages.OrderRequest) error {
	data, err := json.Marshal(order)
	if err != nil {
		return err
	}
	return r.rdb.Publish(r.ctx, "orders.new", data).Err()
}

func (r *RedisClient) PublishCancelOrder(cancel *messages.CancelOrderRequest) error {
	data, err := json.Marshal(cancel)
	if err != nil {
		return err
	}

	return r.rdb.Publish(r.ctx, "orders.cancel", data).Err()
}

func (r *RedisClient) PublishOrderResponse(order *messages.OrderResponse) error {
	data, err := json.Marshal(order)

	if err != nil {
		return err
	}

	return r.rdb.Publish(r.ctx, "orders.response", data).Err()
}

func (r *RedisClient) PublishTrade(trade *messages.TradeEvent) error {
	data, err := json.Marshal(trade)

	if err != nil {
		return err
	}

	return r.rdb.Publish(r.ctx, "trade.executed", data).Err()
}

func (r *RedisClient) PublishOrderBookUpdate(update *messages.OrderBookUpdate) error {
	data, err := json.Marshal(update)
	if err != nil {
		return err
	}
	return r.rdb.Publish(r.ctx, "orderbook.update", data).Err()
}
