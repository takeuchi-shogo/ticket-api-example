package usecase

import "github.com/takeuchi-shogo/ticket-api/internal/domain/models"

type OrganizerUsecase interface {
	FindByID(id int) (*models.Organizers, error)
}
