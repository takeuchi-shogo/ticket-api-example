package usecase

import "github.com/takeuchi-shogo/ticket-api/internal/domain/models"

type StripeUsecase interface {
	GetCustomer(customerID string) (*models.StripeCustomer, error)
	GetCreditCardByCustomerID(customerID, defaultSource string) (*models.StripeCreditCards, error)
	CreateCreditCardAndCustomer(customer *models.StripeCustomer, token string) error
}
