package services

import (
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
)

type TicketService interface {
	Get(id int) (*models.TicketsResponse, *usecase.ResultStatus)
	GetListByEventID(eventID int) (*models.TicketInteractorResponse, *usecase.ResultStatus)
	GetListByArtistID(artistID int) ([]*models.TicketsResponse, *usecase.ResultStatus)
	Create(ticket *models.Tickets) (*models.TicketsResponse, *usecase.ResultStatus)
}
