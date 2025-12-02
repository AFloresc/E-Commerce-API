package payment

import (
	"e-commerce-api/internal/cart"
	"e-commerce-api/internal/product"
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
		// Validar producto contra el servicio de productos
		p, err := product.GetProduct(item.ProductID)
		if err != nil {
			http.Error(w, "Producto inválido en carrito", http.StatusBadRequest)
			return
		}
		if item.Quantity > p.Stock {
			http.Error(w, "Cantidad supera stock disponible", http.StatusBadRequest)
			return
		}

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
