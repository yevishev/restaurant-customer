package models

type Dish struct {
	ID int `json:"dishId"`
	Name string `json:"name"`
	Price string `json:"price"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type DishList struct {
	Dishes []Dish `json:"dishes"` 
}