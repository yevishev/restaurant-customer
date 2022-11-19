package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/yevishev/restaurant-customer/handler"
)

func main() {
	
	var err error = godotenv.Load()

	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	
	http.HandleFunc("/ping", handler.PingHandler)
	http.ListenAndServe(":80", nil)
}