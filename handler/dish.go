package handler

import (
	"fmt"
	"net/http"
)

func CreateMuxDish() *http.ServeMux {
	mux := http.NewServeMux()

    // Регистрируем обработчики для маршрутов в этом модуле
    mux.HandleFunc("/dishes", getAllDishes)
    mux.HandleFunc("/dish", createDish)

	return mux
}
//GetAllOrders
func getAllDishes(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "getAllDishes\n")
}

//CreateOrder
func createDish(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "createDish %s\n", "slug")
}

