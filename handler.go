package handler

import (
	"fmt"
	"net/http"
)

func pingHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "pong")
}