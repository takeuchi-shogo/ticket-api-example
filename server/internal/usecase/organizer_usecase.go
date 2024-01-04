package usecase

import (
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/uptrace/bun"
)

type OrganizerUsecase interface {
	FindByID(db bun.IDB, id int) (*models.Organizers, error)
	Create(db bun.IDB, organizer *models.Organizers) (*models.Organizers, error)
}
