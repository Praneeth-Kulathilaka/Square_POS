package dto

type OrderResponse struct {
	ID        string   	`json:"id"`
	OpenedAt  string   	`json:"opened_at"`
	IsClosed  bool     	`json:"is_closed"`
	Table     string   	`json:"table"`
	Items     []Item   	`json:"items"`
	Totals    Total  	`json:"totals"`
}

type Item struct {
	Name       string         	`json:"name"`
	Comment    string         	`json:"comment"`
	UnitPrice  int64          	`json:"unit_price"`
	Quantity   int            	`json:"quantity"`
	Discounts  []DiscountDTO  	`json:"discounts"`
	Modifiers  []Modifier  		`json:"modifiers"`
	Amount     int64          	`json:"amount"`
}

type DiscountDTO struct {
	Name        string 	`json:"name"`
	IsPercentage bool   `json:"is_percentage"`
	Value       int64  	`json:"value"`
	Amount      int64  	`json:"amount"`
}

type Modifier struct {
	Name      string `json:"name"`
	UnitPrice int64  `json:"unit_price"`
	Quantity  int    `json:"quantity"`
	Amount    int64  `json:"amount"`
}

type Total struct {
	Discounts     int64 `json:"discounts"`
	Due           int64 `json:"due"`
	Tax           int64 `json:"tax"`
	ServiceCharge int64 `json:"service_charge"`
	Paid          int64 `json:"paid"`
	Tips          int64 `json:"tips"`
	Total         int64 `json:"total"`
}