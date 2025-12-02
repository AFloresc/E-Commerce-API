package cart

import (
	"e-commerce-api/middleware"

	"github.com/go-chi/chi/v5"
)

func Router() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.JWTAuth)

	r.Get("/", GetCartHandler)
	r.Post("/add", AddHandler)
	r.Post("/remove", RemoveHandler)

	return r
}
