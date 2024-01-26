package services

import (
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
)

type TicketHasItemService interface {
	Create(ticket *models.TicketHasItems) (*models.TicketHasItemsResponse, *usecase.ResultStatus)
}
