package main

import (
	"net/http"
	"github.com/yevishev/restaurant-customer/pkg/handler"
)

func main() {

	http.HandleFunc("/ping", handler.PingHandler)

	http.ListenAndServe(":80", nil)

}