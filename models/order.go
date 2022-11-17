package models

import "time"

type Order struct {
	ID string `json:"id"`
	UserId string `json:"user_id"`
	Price float64 `json:"price"`
	PaymentStatus bool `json:"payment_status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type OrderList struct {
	Orders []Order `json:"orders"` 
}