package interactors

import (
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase/services"
)

type EventInteractor struct {
	db           usecase.DBUsecase
	eventUsecase usecase.EventUsecase
}

func NewEventInteractor(
	db usecase.DBUsecase,
	event usecase.EventUsecase,
) services.EventService {
	return &EventInteractor{
		db:           db,
		eventUsecase: event,
	}
}

func (e *EventInteractor) Get(id int) (*models.Events, *usecase.ResultStatus) {
	db := e.db.Connect()

	event, err := e.eventUsecase.FindByID(db, id)
	if err != nil {
		return &models.Events{}, usecase.NewResultStatus(400, err)
	}
	return event, usecase.NewResultStatus(200, nil)
}
