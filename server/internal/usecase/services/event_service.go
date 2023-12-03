package services

import (
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
)

type EventService interface {
	Get(id int) (*models.Events, *usecase.ResultStatus)
	GetList() ([]*models.EventsReponse, *usecase.ResultStatus)
	Create(event *models.Events) (*models.EventsReponse, *usecase.ResultStatus)
}
