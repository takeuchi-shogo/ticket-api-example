package usecase

import (
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/uptrace/bun"
)

type TicketItemUsecase interface {
	FindByID(db bun.IDB, id int) (*models.TicketItems, error)
	Create(db bun.IDB, ticket *models.TicketItems) (*models.TicketItems, error)
}
