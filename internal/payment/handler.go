package payment

import (
	"e-commerce-api/internal/cart"
	"e-commerce-api/middleware"
	"encoding/json"
	"net/http"
)

func CheckoutHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey).(string)
	c := cart.GetCart(userID)

	// Convertir carrito a formato Stripe
	items := []map[string]interface{}{}
	for _, item := range c.Items {
		// Aquí deberías obtener datos reales del producto desde product service
		items = append(items, map[string]interface{}{
			"name":     item.ProductID,
			"price":    10.0, // placeholder, reemplazar con precio real
			"quantity": item.Quantity,
		})
	}

	session, err := CreateCheckoutSession(userID, items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"checkout_url": session.URL,
	})
}
