package usecase

import (
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/uptrace/bun"
)

type TicketHasItemUsecase interface {
	FindByID(db bun.IDB, id int) (*models.TicketHasItems, error)
	FindByTicketID(db bun.IDB, ticketID int) ([]*models.TicketHasItems, error)
	FindByTicketItemID(db bun.IDB, ticketItemID int) ([]*models.TicketHasItems, error)
	FindByTicketIDAndTicketItemID(db bun.IDB, ticketID int, ticketItemID int) (*models.TicketHasItems, error)
	Create(db bun.IDB, ticket *models.TicketHasItems) (*models.TicketHasItems, error)
}
