package database

import (
	"context"
	"time"

	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
	"github.com/uptrace/bun"
)

type TicketItemRepository struct{}

func NewTicketItemRepository() usecase.TicketItemUsecase {
	return &TicketItemRepository{}
}

func (t *TicketItemRepository) FindByID(db bun.IDB, id int) (*models.TicketItems, error) {
	ticket := &models.TicketItems{}
	if err := db.NewSelect().
		Model(ticket).
		Where("id = ?", id).
		Scan(context.Background()); err != nil {
		return &models.TicketItems{}, err
	}
	return ticket, nil
}

func (t *TicketItemRepository) Create(db bun.IDB, ticket *models.TicketItems) (*models.TicketItems, error) {

	ticket.CreatedAt = time.Now().Unix()
	ticket.UpdatedAt = time.Now().Unix()

	_, err := db.NewInsert().Model(ticket).Exec(context.Background())

	return ticket, err
}
