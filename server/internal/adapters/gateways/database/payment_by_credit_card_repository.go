package database

import (
	"context"
	"time"

	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
	"github.com/uptrace/bun"
)

type paymentByCreditCardRepository struct{}

func NewPaymentByCreditCardRepository() usecase.PaymentByCreditCardUsecase {
	return &paymentByCreditCardRepository{}
}

func (p *paymentByCreditCardRepository) FindByIsValid(db bun.IDB, isValid bool) ([]*models.PaymentByCreditCards, error) {
	payments := []*models.PaymentByCreditCards{}
	err := db.NewSelect().
		Model(&payments).Where("is_valid = ?", isValid).Scan(context.Background())
	return payments, err
}

func (p *paymentByCreditCardRepository) Create(db bun.IDB, payment *models.PaymentByCreditCards) (*models.PaymentByCreditCards, error) {
	payment.CreatedAt = time.Now().Unix()
	payment.UpdatedAt = time.Now().Unix()

	if _, err := db.NewInsert().Model(payment).Exec(context.Background()); err != nil {
		return &models.PaymentByCreditCards{}, err
	}
	return payment, nil
}

func (p *paymentByCreditCardRepository) Save(db bun.IDB, payment *models.PaymentByCreditCards) (*models.PaymentByCreditCards, error) {
	payment.UpdatedAt = time.Now().Unix()

	_, err := db.NewUpdate().
		Model(payment).
		WherePK().
		Exec(context.Background())

	return payment, err
}