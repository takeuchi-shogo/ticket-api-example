package interactors

import (
	"net/http"

	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase/services"
)

type eventInteractor struct {
	db    usecase.DBUsecase
	event usecase.EventUsecase
}

func NewEventInteractor(
	db usecase.DBUsecase,
	event usecase.EventUsecase,
) services.EventService {
	return &eventInteractor{
		db:    db,
		event: event,
	}
}

func (e *eventInteractor) Get(id int) (*models.Events, *usecase.ResultStatus) {

	db, _ := e.db.Connect()

	event, err := e.event.FindByID(db, id)
	if err != nil {
		return &models.Events{}, usecase.NewResultStatus(400, err)
	}
	return event, usecase.NewResultStatus(200, nil)
}

func (e *eventInteractor) GetList(eventType string) ([]*models.EventsReponse, *usecase.ResultStatus) {

	db, _ := e.db.Connect()

	foundEvents := []*models.Events{}

	if eventType != "" {
		events, err := e.event.FindByEventType(db, eventType)
		if err != nil {
			return []*models.EventsReponse{}, usecase.NewResultStatus(400, err)
		}
		foundEvents = events
	}

	if eventType == "" {
		events, err := e.event.Find(db)
		if err != nil {
			return []*models.EventsReponse{}, usecase.NewResultStatus(400, err)
		}
		foundEvents = events
	}

	builtEvents := []*models.EventsReponse{}

	for _, event := range foundEvents {
		builtEvents = append(builtEvents, event.BuildFor())
	}
	return builtEvents, usecase.NewResultStatus(200, nil)
}

func (e *eventInteractor) GetListByArtistID(eventID, artistID int) (*models.EventInteractorResponse, *usecase.ResultStatus) {

	db, _ := e.db.Connect()

	events, err := e.event.FindByArtistID(db, artistID)
	if err != nil {
		return &models.EventInteractorResponse{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	builtEvents := []*models.EventsReponse{}
	for _, event := range events {
		builtEvents = append(builtEvents, event.BuildFor())
	}
	return &models.EventInteractorResponse{
		Total:  len(builtEvents),
		Events: builtEvents,
	}, usecase.NewResultStatus(200, nil)
}

func (e *eventInteractor) Create(event *models.Events) (*models.EventsReponse, *usecase.ResultStatus) {

	db, _ := e.db.Connect()

	newEvent, err := e.event.Create(db, event)
	if err != nil {
		return &models.EventsReponse{}, usecase.NewResultStatus(400, err)
	}

	return newEvent.BuildFor(), usecase.NewResultStatus(201, nil)
}
