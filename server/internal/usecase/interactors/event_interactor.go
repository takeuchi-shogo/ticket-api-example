package interactors

import (
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase/services"
)

type eventInteractor struct {
	db           usecase.DBUsecase
	eventUsecase usecase.EventUsecase
}

func NewEventInteractor(
	db usecase.DBUsecase,
	event usecase.EventUsecase,
) services.EventService {
	return &eventInteractor{
		db:           db,
		eventUsecase: event,
	}
}

func (e *eventInteractor) Get(id int) (*models.Events, *usecase.ResultStatus) {

	db := e.db.Connect()

	event, err := e.eventUsecase.FindByID(db, id)
	if err != nil {
		return &models.Events{}, usecase.NewResultStatus(400, err)
	}
	return event, usecase.NewResultStatus(200, nil)
}

func (e *eventInteractor) GetList() ([]*models.EventsReponse, *usecase.ResultStatus) {

	db := e.db.Connect()

	events, err := e.eventUsecase.Find(db)
	if err != nil {
		return []*models.EventsReponse{}, usecase.NewResultStatus(400, err)
	}

	builtEvents := []*models.EventsReponse{}

	for _, event := range events {
		builtEvents = append(builtEvents, event.BuildFor())
	}
	return builtEvents, usecase.NewResultStatus(200, nil)
}

func (e *eventInteractor) Create(event *models.Events) (*models.EventsReponse, *usecase.ResultStatus) {

	db := e.db.Connect()

	newEvent, err := e.eventUsecase.Create(db, event)
	if err != nil {
		return &models.EventsReponse{}, usecase.NewResultStatus(400, err)
	}

	return newEvent.BuildFor(), usecase.NewResultStatus(201, nil)
}
