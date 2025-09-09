package types
type OrderResponse struct {
    ID                string `json:"id"`
    MarketID          string `json:"market_id"`
    Side              string `json:"side"`
    Type              string `json:"type"`
    Quantity          string `json:"quantity"`
    Price             string `json:"price,omitempty"`
    FilledQuantity    string `json:"filled_quantity"`
    RemainingQuantity string `json:"remaining_quantity"`
    Status            string `json:"status"`
    CreatedAt         string `json:"created_at"`
}