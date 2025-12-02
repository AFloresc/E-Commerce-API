package payment

import (
	"e-commerce-api/internal/cart"
	"e-commerce-api/middleware"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/stripe/stripe-go/v78"
	stripeSession "github.com/stripe/stripe-go/v78/checkout/session"
	"github.com/stripe/stripe-go/v78/webhook"
)

type Handler struct {
	Service *PaymentService
}

func NewHandler(service *PaymentService) *Handler {
	return &Handler{Service: service}
}

// Checkout protegido con JWT
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

// Webhook público para Stripe
func (h *Handler) WebhookHandler(w http.ResponseWriter, r *http.Request) {
	const MaxBodyBytes = int64(65536)
	r.Body = http.MaxBytesReader(w, r.Body, MaxBodyBytes)
	payload, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error leyendo body", http.StatusServiceUnavailable)
		return
	}

	event, err := webhook.ConstructEvent(
		payload,
		r.Header.Get("Stripe-Signature"),
		os.Getenv("STRIPE_WEBHOOK_SECRET"),
	)
	if err != nil {
		http.Error(w, "Firma inválida", http.StatusBadRequest)
		return
	}

	if event.Type == "checkout.session.completed" {
		var s stripe.CheckoutSession
		if err := json.Unmarshal(event.Data.Raw, &s); err != nil {
			http.Error(w, "Error parseando sesión", http.StatusBadRequest)
			return
		}

		// Recuperar sesión expandida con line_items
		expandedSession, err := stripeSession.Get(s.ID, &stripe.CheckoutSessionParams{
			Expand: []*string{stripe.String("line_items")},
		})
		if err != nil {
			http.Error(w, "Error recuperando line_items", http.StatusInternalServerError)
			return
		}

		items := []cart.CartItem{}
		for _, li := range expandedSession.LineItems.Data {
			// ⚠️ Requiere que hayas guardado product_id en metadata al crear la sesión
			productID := li.Price.Product.ID // idealmente usar metadata["product_id"]
			qty := int(li.Quantity)

			items = append(items, cart.CartItem{
				ProductID: productID,
				Quantity:  qty,
			})
		}

		if err := h.Service.ApplyPurchase(items); err != nil {
			http.Error(w, "Error aplicando compra: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}
