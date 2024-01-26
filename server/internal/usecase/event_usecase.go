package usecase

import (
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/uptrace/bun"
)

type EventUsecase interface {
	CountEventByArtistID(db bun.IDB, artistID int) (int, error)
	Find(db bun.IDB) ([]*models.Events, error)
	FindByID(db bun.IDB, id int) (*models.Events, error)
	FindByArtistID(db bun.IDB, artistID int) ([]*models.Events, error)
	FindByEventType(db bun.IDB, eventType string) ([]*models.Events, error)
	Create(db bun.IDB, event *models.Events) (*models.Events, error)
}
