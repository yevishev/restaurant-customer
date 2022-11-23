package handler

import (
	"fmt"
	"net/http"
	"strconv"
)

var OrderRoutes = []Route{
	NewRoute("GET", "/orders", getAllOrders),
	NewRoute("POST", "/order", createOrder),
	NewRoute("GET", "/order/([^/]+)", order),
	NewRoute("POST", "/order/([^/]+)", updateOrder),
	NewRoute("GET", "/order/([^/]+)/delete", deleteOrder),
}

//GetAllOrders
func getAllOrders(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "apiGetWidgets\n")
}

//CreateOrder
func createOrder(writer http.ResponseWriter, request *http.Request) {
	slug := GetField(request, 0)
	fmt.Fprintf(writer, "createOrder %s\n", slug)
}

//GetOrderById
func order(writer http.ResponseWriter, request *http.Request) {
	id := GetField(request, 0)
	fmt.Fprintf(writer, "order %s\n", id)
}

//DeleteOrder
func deleteOrder(writer http.ResponseWriter, request *http.Request) {
	slug := GetField(request, 0)
	id, _ := strconv.Atoi(GetField(request, 1))
	fmt.Fprintf(writer, "deleteOrder %s %d\n", slug, id)
}

//UpdateOrder
func updateOrder(writer http.ResponseWriter, request *http.Request) {
	slug := GetField(request, 0)
	id, _ := strconv.Atoi(GetField(request, 1))
	fmt.Fprintf(writer, "updateOrder %s %d\n", slug, id)
}
