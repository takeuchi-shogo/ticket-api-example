package database

import (
	"context"
	"time"

	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
	"github.com/uptrace/bun"
)

type userMailLogRepository struct{}

func NewUserMailLogRepository() usecase.UserMailLogUsecase {
	return &userMailLogRepository{}
}

func (u *userMailLogRepository) Create(db bun.IDB, userMailLog *models.UserMailLogs) (*models.UserMailLogs, error) {

	userMailLog.CreatedAt = time.Now().Unix()
	userMailLog.UpdatedAt = time.Now().Unix()

	_, err := db.NewInsert().Model(userMailLog).Exec(context.Background())

	return userMailLog, err
}
