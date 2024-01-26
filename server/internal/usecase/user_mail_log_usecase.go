package usecase

import (
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/uptrace/bun"
)

type UserMailLogUsecase interface {
	Create(db bun.IDB, userMailLog *models.UserMailLogs) (*models.UserMailLogs, error)
}
