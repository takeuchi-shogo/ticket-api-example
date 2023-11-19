package database

import (
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
)

type eventRepository struct{}

func NewEventRepository() usecase.EventUsecase {
	return &eventRepository{}
}

func (e *eventRepository) FindByID(id int) (*models.Events, error) {
	return &models.Events{}, nil
}
