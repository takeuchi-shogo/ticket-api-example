package usecase

import (
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/uptrace/bun"
)

type UserPaymentHistoryUsecase interface {
	Create(db bun.IDB, userPaymentHistory *models.UserPaymentHistories) (*models.UserPaymentHistories, error)
}
