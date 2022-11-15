package handler

import (
	"fmt"
	"net/http"
)

func PingHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "pong")
}