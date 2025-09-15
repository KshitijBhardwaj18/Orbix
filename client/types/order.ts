export interface PlaceOrderResponse {
  id: string;
  market_id: string;
  side: "BUY" | "SELL";
  type: "MARKET" | "LIMIT";
  quantity: string;
  price?: string;
  filledQuantity: string;
  remaining_quantity: string;
  status: "PENDING" | "FILLED" | "PARTIAL" | "CANCELLED" | "REJECTED";
  created_at: string;
}

export interface PlaceOrderRequest {
  "market-id": string;
  side: string;
  type: string;
  quantity: string;
  price: string;
}

type orderType = "LIMIT" | "MARKET";
type orderStatus = "PENDING" | "FILLED" | "PARTIAL" | "CANCELLED" | "REJECTED";
type orderSide = "BUY" | "SELL";
