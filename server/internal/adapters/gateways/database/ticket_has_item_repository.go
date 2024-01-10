package database

import (
	"context"
	"time"

	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
	"github.com/uptrace/bun"
)

type TicketHasItemRepository struct{}

func NewTicketHasItemRepository() usecase.TicketHasItemUsecase {
	return &TicketHasItemRepository{}
}

func (t *TicketHasItemRepository) FindByID(db bun.IDB, id int) (*models.TicketHasItems, error) {
	ticket := &models.TicketHasItems{}
	if err := db.NewSelect().
		Model(ticket).
		Where("id = ?", id).
		Scan(context.Background()); err != nil {
		return &models.TicketHasItems{}, err
	}
	return ticket, nil
}

func (t *TicketHasItemRepository) FindByTicketID(db bun.IDB, ticketID int) ([]*models.TicketHasItems, error) {
	tickets := []*models.TicketHasItems{}
	if err := db.NewSelect().
		Model(&tickets).
		Where("ticket_id = ?", ticketID).
		Scan(context.Background()); err != nil {
		return []*models.TicketHasItems{}, err
	}
	return tickets, nil
}

func (t *TicketHasItemRepository) FindByTicketItemID(db bun.IDB, ticketItemID int) ([]*models.TicketHasItems, error) {
	tickets := []*models.TicketHasItems{}
	if err := db.NewSelect().
		Model(tickets).
		Where("ticket_item_id = ?", ticketItemID).
		Scan(context.Background()); err != nil {
		return []*models.TicketHasItems{}, err
	}
	return tickets, nil
}

func (t *TicketHasItemRepository) FindByTicketIDAndTicketItemID(db bun.IDB, ticketID int, ticketItemID int) (*models.TicketHasItems, error) {
	ticket := &models.TicketHasItems{}
	if err := db.NewSelect().
		Model(ticket).
		Where("ticket_id = ? and ticket_item_id", ticketID, ticketItemID).
		Scan(context.Background()); err != nil {
		return &models.TicketHasItems{}, err
	}
	return ticket, nil
}

func (t *TicketHasItemRepository) Create(db bun.IDB, ticket *models.TicketHasItems) (*models.TicketHasItems, error) {

	ticket.CreatedAt = time.Now().Unix()
	ticket.UpdatedAt = time.Now().Unix()

	_, err := db.NewInsert().Model(ticket).Exec(context.Background())
	if err != nil {
		return &models.TicketHasItems{}, err
	}

	return ticket, err
}
