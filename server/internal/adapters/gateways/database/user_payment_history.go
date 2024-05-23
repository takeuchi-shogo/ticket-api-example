package database

import (
	"context"
	"time"

	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
	"github.com/uptrace/bun"
)

type UserPaymentHistoryRepository struct{}

func NewUserPaymentHistoryRepository() usecase.UserPaymentHistoryUsecase {
	return &UserPaymentHistoryRepository{}
}

func (u *UserPaymentHistoryRepository) Create(db bun.IDB, userPaymentHistory *models.UserPaymentHistories) (*models.UserPaymentHistories, error) {
	ctx := context.Background()

	userPaymentHistory.CreatedAt = time.Now().Unix()

	_, err := db.NewInsert().Model(userPaymentHistory).Exec(ctx)
	if err != nil {
		return &models.UserPaymentHistories{}, err
	}
	return userPaymentHistory, nil
}
