package orders

import (
	"encoding/json"
	"log"
)

func IngestOrder(data []byte) (*Order, error) {
	var order Order
	if err := json.Unmarshal(data, &order); err != nil {
		log.Printf("Failed to parse order: %v", err)
		return nil, err
	}
	order.Status = "Received"
	return &order, nil
}
