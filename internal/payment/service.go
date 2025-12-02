package payment

import (
	"os"

	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/checkout/session"
)

func InitStripe() {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
}

func CreateCheckoutSession(userID string, items []map[string]interface{}) (*stripe.CheckoutSession, error) {
	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		Mode:               stripe.String("payment"),
		SuccessURL:         stripe.String("http://localhost:8080/success"),
		CancelURL:          stripe.String("http://localhost:8080/cancel"),
	}

	// Convertir items del carrito a Stripe line items
	for _, item := range items {
		params.LineItems = append(params.LineItems, &stripe.CheckoutSessionLineItemParams{
			PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
				Currency: stripe.String("eur"),
				ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
					Name: stripe.String(item["name"].(string)),
				},
				UnitAmount: stripe.Int64(int64(item["price"].(float64) * 100)), // en céntimos
			},
			Quantity: stripe.Int64(int64(item["quantity"].(int))),
		})
	}

	return session.New(params)
}
