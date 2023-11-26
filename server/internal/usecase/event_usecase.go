package usecase

import (
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/uptrace/bun"
)

type EventUsecase interface {
	FindByID(db *bun.DB, id int) (*models.Events, error)
}
