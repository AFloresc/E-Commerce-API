package cart

import (
	"e-commerce-api/middleware"
	"encoding/json"
	"net/http"
)

func GetCartHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey).(string)
	c := GetCart(userID)
	json.NewEncoder(w).Encode(c)
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey).(string)
	var req struct {
		ProductID string `json:"product_id"`
		Quantity  int    `json:"quantity"`
	}
	json.NewDecoder(r.Body).Decode(&req)
	c := AddToCart(userID, req.ProductID, req.Quantity)
	json.NewEncoder(w).Encode(c)
}

func RemoveHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey).(string)
	var req struct {
		ProductID string `json:"product_id"`
	}
	json.NewDecoder(r.Body).Decode(&req)
	c, err := RemoveFromCart(userID, req.ProductID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(c)
}
