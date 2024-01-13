package database

import (
	"context"

	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
	"github.com/uptrace/bun"
)

type creditCardRepository struct{}

func NewCreditCardRepository() usecase.CreditCardUsecase {
	return &creditCardRepository{}
}

func (c *creditCardRepository) FindByUserID(db bun.IDB, userID int) (*models.CreditCards, error) {
	creditCard := &models.CreditCards{}
	if err := db.NewSelect().
		Model(creditCard).
		Where("user_id = ?", userID).
		Scan(context.Background()); err != nil {
		return &models.CreditCards{}, err
	}
	return creditCard, nil
}
