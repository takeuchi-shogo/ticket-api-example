package interactors

import (
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase/services"
)

type EventInteractor struct {
	eventUsecase usecase.EventUsecase
}

func NewEventInteractor(
	event usecase.EventUsecase,
) services.EventService {
	return &EventInteractor{
		eventUsecase: event,
	}
}

func (e *EventInteractor) Get(id int) (*models.Events, *usecase.ResultStatus) {
	event, err := e.eventUsecase.FindByID(id)
	if err != nil {
		return &models.Events{}, usecase.NewResultStatus(400, err)
	}
	return event, usecase.NewResultStatus(200, nil)
}
