package usecase

import (
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/uptrace/bun"
)

type EventUsecase interface {
	Find(db bun.IDB) ([]*models.Events, error)
	FindByID(db bun.IDB, id int) (*models.Events, error)
	Create(db bun.IDB, event *models.Events) (*models.Events, error)
}
