package database

import (
	"context"

	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
	"github.com/uptrace/bun"
)

type paymentByCreditCardRepository struct{}

func NewPaymentByCreditCardRepository() usecase.PaymentByCreditCardUsecase {
	return &paymentByCreditCardRepository{}
}

func (p *paymentByCreditCardRepository) Create(db bun.IDB, payment *models.PaymentByCreditCards) (*models.PaymentByCreditCards, error) {
	if _, err := db.NewInsert().Model(payment).Exec(context.Background()); err != nil {
		return &models.PaymentByCreditCards{}, err
	}
	return payment, nil
}
