package payment

import (
	"e-commerce-api/middleware"

	"github.com/go-chi/chi/v5"
)

func Router(service *PaymentService) chi.Router {
	h := NewHandler(service)
	r := chi.NewRouter()
	r.Use(middleware.JWTAuth)

	r.Post("/checkout", h.CheckoutHandler)

	return r
}
