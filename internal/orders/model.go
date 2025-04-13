package orders

type Order struct {
	ID     string `json:"id"`
	Source string `json:"source"`
	Status string `json:"status"`
	Items  []string `json:"items"`
}
