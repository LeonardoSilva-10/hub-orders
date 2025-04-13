package main

import (
	"hub-orders/internal/api"
	"hub-orders/internal/db"
	"hub-orders/internal/logger"
	"log"
	"net/http"
)

func main() {
	logger.Init()
	db.Init()

	http.HandleFunc("/orders", api.GetOrdersHandler)
	http.HandleFunc("/orders/update", api.UpdateOrderHandler)
	http.HandleFunc("/orders/ingest", api.IngestOrderHandler)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
