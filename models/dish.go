package models

type Dish struct {
	ID int `json:"dishId"`
	Name string `json:"name"`
	Price string `json:"price"`
}

type DishList struct {
	Dishes []Dish `json:"dishes"` 
}