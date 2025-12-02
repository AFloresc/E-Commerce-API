package payment

import (
	"e-commerce-api/internal/cart"
	"e-commerce-api/middleware"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/webhook"
)

type Handler struct {
	Service *PaymentService
}

func NewHandler(service *PaymentService) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) CheckoutHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey).(string)
	c := cart.GetCart(userID)
	session, err := h.Service.CreateCheckoutSession(userID, c.Items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"checkout_url": session.URL,
	})
}

func (h *Handler) WebhookHandler(w http.ResponseWriter, r *http.Request) {
	const MaxBodyBytes = int64(65536)
	r.Body = http.MaxBytesReader(w, r.Body, MaxBodyBytes)
	payload, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error leyendo body", http.StatusServiceUnavailable)
		return
	}

	event, err := webhook.ConstructEvent(payload, r.Header.Get("Stripe-Signature"), os.Getenv("STRIPE_WEBHOOK_SECRET"))
	if err != nil {
		http.Error(w, "Firma inválida", http.StatusBadRequest)
		return
	}

	// Procesar evento de checkout completado
	if event.Type == "checkout.session.completed" {
		var session stripe.CheckoutSession
		if err := json.Unmarshal(event.Data.Raw, &session); err != nil {
			http.Error(w, "Error parseando sesión", http.StatusBadRequest)
			return
		}

		// Aquí deberías recuperar los items comprados.
		// En producción, se usan `session.Metadata` o `line_items` para saber qué productos se compraron.
		// Ejemplo simplificado: simulamos items desde metadata
		items := []cart.CartItem{
			{ProductID: "example-product-id", Quantity: 1},
		}

		// Reducir stock usando PaymentService
		if err := h.Service.ApplyPurchase(items); err != nil {
			http.Error(w, "Error aplicando compra: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}
