package usecase

import "github.com/takeuchi-shogo/ticket-api/internal/domain/models"

type TicketUsecase interface {
	FindByID(id int) (*models.Tickets, error)
}
