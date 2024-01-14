package services

import (
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
)

type UserBookTicketService interface {
	Get(bookID string) (*models.UserBookTicketsReponse, *usecase.ResultStatus)
}
