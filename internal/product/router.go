package product

import (
	"e-commerce-api/middleware"

	"github.com/go-chi/chi/v5"
)

func Router(repo *ProductRepo) chi.Router {
	h := NewHandler(repo)
	r := chi.NewRouter()

	// Rutas públicas
	r.Get("/", h.ListHandler)
	r.Get("/{id}", h.GetHandler)

	// Rutas protegidas (solo admin)
	r.Group(func(r chi.Router) {
		r.Use(middleware.JWTAuth)
		r.Use(middleware.RequireAdmin)
		r.Post("/", h.CreateHandler)
		r.Put("/{id}", h.UpdateHandler)
		r.Delete("/{id}", h.DeleteHandler)
	})

	return r
}
