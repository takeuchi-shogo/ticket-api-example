package usecase

import (
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/uptrace/bun"
)

type PaymentByCreditCardUsecase interface {
	Create(db bun.IDB, payment *models.PaymentByCreditCards) (*models.PaymentByCreditCards, error)
}
