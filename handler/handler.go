package handler

import (
	"database/sql"
	"log"

	"fmt"
	_ "github.com/lib/pq"
	"net/http"
)

func PingHandler(writer http.ResponseWriter, request *http.Request) {
	connStr := "user=pqgotest dbname=pqgotest sslmode=verify-full"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	age := 21
	rows, err := db.Query("SELECT name FROM users WHERE age = $1", age)
	fmt.Print(rows)
	fmt.Fprintln(writer, "pong")
}
