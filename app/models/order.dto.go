package models

//Data types for creating order
type OrderRequest struct {
	IdempotencyKey string `json:"idempotency_key"`
	Order          Order  `json:"order"`
}

type Order struct {
	LocationID string       `json:"location_id"`
	LineItems  []LineItem   `json:"line_items"`
}

type LineItem struct {
	Name          string `json:"name"`
	Quantity      string `json:"quantity"`
	BasePriceMoney struct {
		Amount   int    `json:"amount"`
		Currency string `json:"currency"`
	} `json:"base_price_money"`
}


