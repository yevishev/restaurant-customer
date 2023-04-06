package handler

import (
	"fmt"
	"io"
	"net/http"

)
func CreateMuxOrder() *http.ServeMux {
	mux := http.NewServeMux()

    // Регистрируем обработчики для маршрутов в этом модуле
    mux.HandleFunc("/orders", getAllOrders)
    mux.HandleFunc("/order", createOrder)
    mux.HandleFunc("/order/([^/]+)", order)
    mux.HandleFunc("/order", updateOrder)
    mux.HandleFunc("/order/([^/]+)/delete", deleteOrder)
	
	return mux
}
//GetAllOrders
func getAllOrders(writer http.ResponseWriter, request *http.Request) {
    resp, err := http.Get("https://google.com/")
    if err != nil {
        // handle error
    }
    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)
	fmt.Fprint(writer, string(body))
}

//CreateOrder
func createOrder(writer http.ResponseWriter, request *http.Request) {
    fmt.Fprintf(writer, "apiUpdateWidgetPart %s %d\n", "slug", "id")
}

//GetOrderById
func order(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "order %s\n", "id")
}

//DeleteOrder
func deleteOrder(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "deleteOrder %s %d\n", "slug", "id")

}

//UpdateOrder
func updateOrder(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "updateOrder %s %d\n", "slug", "id")
}
