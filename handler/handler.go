package handler

import (
	"fmt"
	"net/http"
)

func CreateMux() *http.ServeMux {
	// Создаем новый ServeMux
	mux := http.NewServeMux()

	// Импортируем обработчики из каждого модуля и регистрируем их в объекте ServeMux
	mux.Handle("/", CreateMuxDish())
	mux.Handle("/", CreateMuxOrder())
	mux.Handle("/test", CreateMuxOrder())

	return mux
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
