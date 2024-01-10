package services

import (
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
)

type TicketItemService interface {
	Create(ticket *models.TicketItems) (*models.TicketItemsResponse, *usecase.ResultStatus)
}
