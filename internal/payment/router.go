package payment

import (
	"e-commerce-api/middleware"

	"github.com/go-chi/chi/v5"
)

func Router() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.JWTAuth)

	r.Post("/checkout", CheckoutHandler)

	return r
}
