package usecase

import (
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/uptrace/bun"
)

type TicketUsecase interface {
	FindByID(db bun.IDB, id int) (*models.Tickets, error)
	FindByEventID(db bun.IDB, eventID int) ([]*models.Tickets, error)
	FindByArtistID(db bun.IDB, artistID int) ([]*models.Tickets, error)
	FindByDrawingAt(db bun.IDB, currentTiem int64) ([]*models.Tickets, error)
	Create(db bun.IDB, ticket *models.Tickets) (*models.Tickets, error)
}
