package payment

import (
	"e-commerce-api/internal/cart"
	"e-commerce-api/middleware"
	"encoding/json"
	"net/http"
)

type Handler struct {
	Service *PaymentService
}

func NewHandler(service *PaymentService) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) CheckoutHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey).(string)
	c := cart.GetCart(userID)
	session, err := h.Service.CreateCheckoutSession(userID, c.Items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"checkout_url": session.URL,
	})
}
