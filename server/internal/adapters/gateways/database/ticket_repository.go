package database

import (
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
)

type TicketRepository struct {
}

func NewTicketRepository() usecase.TicketUsecase {
	return &TicketRepository{}
}

func (t *TicketRepository) FindByID(id int) (*models.Tickets, error) {
	return &models.Tickets{}, nil
}
