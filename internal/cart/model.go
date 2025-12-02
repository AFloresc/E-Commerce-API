package cart

type CartItem struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type Cart struct {
	UserID string     `json:"user_id"`
	Items  []CartItem `json:"items"`
}
