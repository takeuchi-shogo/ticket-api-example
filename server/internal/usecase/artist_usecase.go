package usecase

import (
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/uptrace/bun"
)

type ArtistUsecase interface {
	Find(db bun.IDB) ([]*models.Artists, error)
	FindByID(db bun.IDB, id int) (*models.Artists, error)
	Create(db bun.IDB, artist *models.Artists) (*models.Artists, error)
}
