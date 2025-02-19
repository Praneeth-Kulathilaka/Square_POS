package models

// Data types for making payment request
type PaymentRequest struct {
	IdempotencyKey string `json:"idempotency_key"`
	AmountMoney struct {
		Amount int `json:"amount"`
		Currency string `json:"currency"`
	}`json:"amount_money"`
	SourceID string `json:"source_id"`
	ReferenceId string `json:"reference_id"`
	OrderId string `json:"order_id"`
}