package usecase

import (
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/uptrace/bun"
)

type UserHasTicketUsecase interface {
	Create(db bun.IDB, userHasTicket *models.UserHasTickets) (*models.UserHasTickets, error)
}
