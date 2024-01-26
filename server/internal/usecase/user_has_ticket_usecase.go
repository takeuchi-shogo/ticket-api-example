package usecase

import (
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/uptrace/bun"
)

type UserHasTicketUsecase interface {
	FindByUserBookTicketID(db bun.IDB, userBookTicketID int) ([]*models.UserHasTickets, error)
	FindByTicketIDAndTicketStatus(db bun.IDB, ticketID int, ticketStatus []string) ([]*models.UserHasTickets, error)
	Create(db bun.IDB, userHasTicket *models.UserHasTickets) (*models.UserHasTickets, error)
	Save(db bun.IDB, userHasTicket *models.UserHasTickets) (*models.UserHasTickets, error)
}
