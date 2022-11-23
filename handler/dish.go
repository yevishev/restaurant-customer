package handler

import (
	"fmt"
	"net/http"
)

var DishRoutes = []Route{
	NewRoute("GET", "/dishes", getAllDishes),
	NewRoute("POST", "/dish", createDish),
}

//GetAllOrders
func getAllDishes(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "getAllDishes\n")
}

//CreateOrder
func createDish(writer http.ResponseWriter, request *http.Request) {
	slug := GetField(request, 0)
	fmt.Fprintf(writer, "createDish %s\n", slug)
}

