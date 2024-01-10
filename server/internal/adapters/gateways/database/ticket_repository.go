package database

import (
	"context"
	"time"

	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
	"github.com/uptrace/bun"
)

type TicketRepository struct {
}

func NewTicketRepository() usecase.TicketUsecase {
	return &TicketRepository{}
}

func (t *TicketRepository) FindByID(db bun.IDB, id int) (*models.Tickets, error) {
	ticket := &models.Tickets{}
	if err := db.NewSelect().
		Model(ticket).
		Where("id = ?", id).
		Scan(context.Background()); err != nil {
		return &models.Tickets{}, err
	}
	return ticket, nil
}

func (t *TicketRepository) Create(db bun.IDB, ticket *models.Tickets) (*models.Tickets, error) {

	ticket.CreatedAt = time.Now().Unix()
	ticket.UpdatedAt = time.Now().Unix()

	_, err := db.NewInsert().Model(ticket).Exec(context.Background())
	if err != nil {
		return &models.Tickets{}, err
	}

	return ticket, err
}
