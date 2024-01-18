package usecase

import (
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/uptrace/bun"
)

type PaymentByCreditCardUsecase interface {
	FindByIsValid(db bun.IDB, isValid bool) ([]*models.PaymentByCreditCards, error)
	Create(db bun.IDB, payment *models.PaymentByCreditCards) (*models.PaymentByCreditCards, error)
	Save(db bun.IDB, payment *models.PaymentByCreditCards) (*models.PaymentByCreditCards, error)
}
