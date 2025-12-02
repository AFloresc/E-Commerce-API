package product

import (
	"e-commerce-api/middleware"

	"github.com/go-chi/chi/v5"
)

func Router() chi.Router {
	r := chi.NewRouter()

	// Rutas públicas
	r.Get("/", ListHandler)
	r.Get("/{id}", GetHandler)

	// Rutas protegidas (solo admin)
	r.Group(func(r chi.Router) {
		r.Use(middleware.JWTAuth)
		r.Use(middleware.RequireAdmin)
		r.Post("/", CreateHandler)
		r.Put("/{id}", UpdateHandler)
		r.Delete("/{id}", DeleteHandler)
	})

	return r
}
