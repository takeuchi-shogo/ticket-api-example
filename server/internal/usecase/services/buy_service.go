package services

import (
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
)

type BuyService interface {
	Create(userBookTicket *models.UserBookTickets) (*models.UserBookTicketsReponse, *usecase.ResultStatus)
}
