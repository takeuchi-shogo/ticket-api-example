package usecase

import (
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/uptrace/bun"
)

type UserBookTicketUsecase interface {
	FindByID(db bun.IDB, id int) (*models.UserBookTickets, error)
	FindByBookID(db bun.IDB, bookID string) (*models.UserBookTickets, error)
	FindByUserID(db bun.IDB, userID int) ([]*models.UserBookTickets, error)
	Create(db bun.IDB, userBookTicket *models.UserBookTickets) (*models.UserBookTickets, error)
}
