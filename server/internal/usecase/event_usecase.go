package usecase

import (
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/uptrace/bun"
)

type EventUsecase interface {
	Find(db *bun.DB) ([]*models.Events, error)
	FindByID(db *bun.DB, id int) (*models.Events, error)
	Create(db *bun.DB, event *models.Events) (*models.Events, error)
}
