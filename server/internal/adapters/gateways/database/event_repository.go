package database

import (
	"errors"

	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
	"github.com/uptrace/bun"
)

type eventRepository struct{}

func NewEventRepository() usecase.EventUsecase {
	return &eventRepository{}
}

func (e *eventRepository) FindByID(db *bun.DB, id int) (*models.Events, error) {
	event := &models.Events{}
	_ = db.NewSelect().Model(event).Where("id = ?")
	if event.ID <= 0 {
		return &models.Events{}, errors.New("event is not found")
	}
	return &models.Events{}, nil
}
