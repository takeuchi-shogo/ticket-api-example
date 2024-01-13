package usecase

import (
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/uptrace/bun"
)

type CreditCardUsecase interface {
	FindByUserID(db bun.IDB, userID int) (*models.CreditCards, error)
}
