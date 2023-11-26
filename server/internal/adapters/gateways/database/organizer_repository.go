package database

import (
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
)

type organizerRepository struct{}

func NewOrganizerRepository() usecase.OrganizerUsecase {
	return &organizerRepository{}
}

func (o *organizerRepository) FindByID(id int) (*models.Organizers, error) {
	return &models.Organizers{}, nil
}
