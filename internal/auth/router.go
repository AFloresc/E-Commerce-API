package auth

import "github.com/go-chi/chi/v5"

func Router() chi.Router {
	r := chi.NewRouter()
	r.Post("/signup", SignupHandler)
	r.Post("/login", LoginHandler)
	return r
}
