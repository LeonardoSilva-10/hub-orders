package api

import (
	"encoding/json"
	"net/http"

	"github.com/hub-orders/internal/db"
	"github.com/hub-orders/internal/orders"
)

var orderStore = make(map[string]*orders.Order)

func GetOrdersHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(orderStore)
}

func UpdateOrderHandler(w http.ResponseWriter, r *http.Request) {
	var order orders.Order
	json.NewDecoder(r.Body).Decode(&order)
	if o, ok := orderStore[order.ID]; ok {
		orders.UpdateOrderStatus(o, order.Status)
		json.NewEncoder(w).Encode(o)
	} else {
		http.Error(w, "Order not found", http.StatusNotFound)
	}
}

func IngestOrderHandler(w http.ResponseWriter, r *http.Request) {
	var order orders.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	_, err := db.DB.Exec("INSERT INTO orders (id, source, status) VALUES ($1, $2, $3)", order.ID, order.Source, order.Status)
	if err != nil {
		http.Error(w, "Failed to insert order", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(order)
}
