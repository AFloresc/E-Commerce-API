package payment

import (
	"e-commerce-api/internal/cart"
	"e-commerce-api/internal/product"
	"os"

	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/checkout/session"
)

type PaymentService struct {
	Repo *product.ProductRepo
}

func NewPaymentService(repo *product.ProductRepo) *PaymentService {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
	return &PaymentService{Repo: repo}
}

// Crear sesión de checkout validando productos contra ProductRepo
func (ps *PaymentService) CreateCheckoutSession(userID string, items []cart.CartItem) (*stripe.CheckoutSession, error) {
	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		Mode:               stripe.String("payment"),
		SuccessURL:         stripe.String("http://localhost:8080/success"),
		CancelURL:          stripe.String("http://localhost:8080/cancel"),
	}

	for _, item := range items {
		// Validar producto contra el repositorio
		p, err := ps.Repo.Get(item.ProductID)
		if err != nil {
			return nil, err
		}
		if item.Quantity > p.Stock {
			return nil, err
		}

		params.LineItems = append(params.LineItems, &stripe.CheckoutSessionLineItemParams{
			PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
				Currency: stripe.String("eur"),
				ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
					Name: stripe.String(p.Name),
				},
				UnitAmount: stripe.Int64(int64(p.Price * 100)), // céntimos
			},
			Quantity: stripe.Int64(int64(item.Quantity)),
		})
	}

	return session.New(params)
}

// Reducir stock tras pago exitoso
func (ps *PaymentService) ApplyPurchase(items []cart.CartItem) error {
	for _, item := range items {
		if err := ps.Repo.ReduceStock(item.ProductID, item.Quantity); err != nil {
			return err
		}
	}
	return nil
}
