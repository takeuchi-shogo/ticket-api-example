package gateways

import (
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
)

type Stripe interface {
	AuthenticatePaymentIntent(amount int, customerID string) (string, error)
	CapturePaymentIntent(paymentID string) (*models.StripePaymentIntents, error)
	GetCustomer(customerID string) (*models.StripeCustomer, error)
	GetCreditCardByCustomerID(customerID, defaultSource string) (*models.StripeCreditCards, error)
	CreateCreditCardAndCustomer(customer *models.StripeCustomer, token string) error
}

type StripeGateway struct {
	stripe Stripe
}

func NewStripeGateway(stripe Stripe) usecase.StripeUsecase {
	return &StripeGateway{
		stripe: stripe,
	}
}

func (s *StripeGateway) GetCustomer(customerID string) (*models.StripeCustomer, error) {
	return s.stripe.GetCustomer(customerID)
}

func (s *StripeGateway) GetCreditCardByCustomerID(customerID, defaultSource string) (*models.StripeCreditCards, error) {
	return s.stripe.GetCreditCardByCustomerID(customerID, defaultSource)
}

func (s *StripeGateway) CreateCreditCardAndCustomer(customer *models.StripeCustomer, token string) error {
	return s.stripe.CreateCreditCardAndCustomer(customer, token)
}

func (s *StripeGateway) AuthenticatePaymentIntent(amount int, customerID string) (string, error) {
	return s.stripe.AuthenticatePaymentIntent(amount, customerID)
}

func (s *StripeGateway) CapturePaymentIntent(paymentID string) (*models.StripePaymentIntents, error) {
	return s.stripe.CapturePaymentIntent(paymentID)
}
