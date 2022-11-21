package models

type Order struct {
	ID int `json:"id"`
	UsersId string `json:"users_id"`
	Price float64 `json:"price"`
	PaymentStatus bool `json:"payment_status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type OrderList struct {
	Orders []Order `json:"orders"` 
}