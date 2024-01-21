package services

import (
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
)

type EventService interface {
	Get(id int) (*models.Events, *usecase.ResultStatus)
	GetList(eventType string) ([]*models.EventsReponse, *usecase.ResultStatus)
	GetListByArtistID(eventID, artistID int) (*models.EventInteractorResponse, *usecase.ResultStatus)
	Create(event *models.Events) (*models.EventsReponse, *usecase.ResultStatus)
}
