package orders

import "log"

func ValidateOrder(order *Order) bool {
	if order.ID == "" || len(order.Items) == 0 {
		log.Printf("Validation failed for order ID %s", order.ID)
		return false
	}
	return true
}

func UpdateOrderStatus(order *Order, status string) {
	log.Printf("Updating order %s to status: %s", order.ID, status)
	order.Status = status
}
