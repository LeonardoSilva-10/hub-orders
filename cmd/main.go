package main

import (
	"log"
	"net/http"

	"github.com/hub-orders/internal/api"
	"github.com/hub-orders/internal/db"
	"github.com/hub-orders/internal/logger"
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
