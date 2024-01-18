package database

import (
	"context"
	"time"

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

func (c *creditCardRepository) Create(db bun.IDB, creditCard *models.CreditCards) (*models.CreditCards, error) {

	creditCard.CreatedAt = time.Now().Unix()
	creditCard.UpdatedAt = time.Now().Unix()

	_, err := db.NewInsert().Model(creditCard).Exec(context.Background())
	if err != nil {
		return &models.CreditCards{}, err
	}
	return creditCard, nil
}
