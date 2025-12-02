package payment

import (
	"e-commerce-api/middleware"

	"github.com/go-chi/chi/v5"
)

func Router(service *PaymentService) chi.Router {
	h := NewHandler(service)
	r := chi.NewRouter()

	// Checkout protegido
	r.Group(func(r chi.Router) {
		r.Use(middleware.JWTAuth)
		r.Post("/checkout", h.CheckoutHandler)
	})

	// Webhook público (Stripe necesita acceso sin JWT)
	r.Post("/webhook", h.WebhookHandler)

	return r
}
