package usecase

import (
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/uptrace/bun"
)

type TicketUsecase interface {
	FindByID(db bun.IDB, id int) (*models.Tickets, error)
	Create(db bun.IDB, ticket *models.Tickets) (*models.Tickets, error)
}
