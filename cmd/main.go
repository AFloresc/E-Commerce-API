package main

import (
	"e-commerce-api/internal/auth"
	"e-commerce-api/internal/cart"
	"e-commerce-api/internal/payment"
	"e-commerce-api/internal/product"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Cargar configuración (puerto, claves, etc.)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Inicializar router
	r := chi.NewRouter()

	// Middlewares básicos
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * 1e9)) // 60 segundos

	// Rutas públicas
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Bienvenido a la E-Commerce API 🚀"))
	})

	// Aquí se montarán los routers de cada módulo
	// ej: r.Mount("/auth", auth.Router())
	//     r.Mount("/products", product.Router())
	r.Mount("/auth", auth.Router())
	r.Mount("/products", product.Router())
	r.Mount("/cart", cart.Router())
	payment.InitStripe()
	r.Mount("/payment", payment.Router())

	log.Printf("Servidor escuchando en http://localhost:%s", port)
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal(err)
	}
}
