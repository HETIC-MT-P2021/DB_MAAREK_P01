package models

import "time"

type PublicOrder struct {
	Number      uint64  `json:"order_number"`
	ItemsNumber uint64  `json:"items_number"`
	TotalPrice  float64 `json:"total_price"`
}

type Order struct {
	Number            uint64    `json:"order_number"`
	Date              time.Time `json:"date"`
	CustomerFirstName string    `json:"customer_first_name"`
	CustomerLastName  string    `json:"customer_lastName"`
	ItemsNumber       uint64    `json:"items_number"`
	TotalPrice        float64   `json:"total_price"`
}

type OrderDetails struct {
	ProductCode        string  `json:"product_code"`
	ProductName        string  `json:"product_name"`
	ProductDescription string  `json:"product_description"`
	Quantity           uint64  `json:"quantity"`
	UnitPrice          float64 `json:"unit_price"`
	LinePrice          float64 `json:"line_price"`
}
