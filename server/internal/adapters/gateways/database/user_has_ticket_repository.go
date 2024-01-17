package database

import (
	"context"
	"time"

	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
	"github.com/uptrace/bun"
)

type userHasTicketRepository struct{}

func NewUserHasTicketRepository() usecase.UserHasTicketUsecase {
	return &userHasTicketRepository{}
}

func (u *userHasTicketRepository) FindByUserBookTicketID(db bun.IDB, userBookTicketID int) ([]*models.UserHasTickets, error) {
	userHasTickets := []*models.UserHasTickets{}
	if err := db.NewSelect().
		Model(userHasTickets).
		Where("user_book_ticket_id = ?", userBookTicketID).
		Scan(context.Background()); err != nil {
		return []*models.UserHasTickets{}, err
	}
	return userHasTickets, nil
}

func (u *userHasTicketRepository) FindByTicketIDAndTicketStatus(db bun.IDB, ticketID int, ticketStatus []string) ([]*models.UserHasTickets, error) {
	userHasTickets := []*models.UserHasTickets{}
	if err := db.NewSelect().
		Model(userHasTickets).
		Where("ticket_id = ?", ticketID).
		Where("ticket_stauts IN (?)", bun.In(ticketStatus)).
		Scan(context.Background()); err != nil {
		return []*models.UserHasTickets{}, err
	}
	return userHasTickets, nil
}

func (u *userHasTicketRepository) Create(db bun.IDB, userHasTicket *models.UserHasTickets) (*models.UserHasTickets, error) {
	_, err := db.NewInsert().Model(userHasTicket).Exec(context.Background())
	return userHasTicket, err
}

func (u *userHasTicketRepository) Save(db bun.IDB, userHasTicket *models.UserHasTickets) (*models.UserHasTickets, error) {
	userHasTicket.UpdatedAt = time.Now().Unix()

	_, err := db.NewUpdate().
		Model(userHasTicket).
		WherePK().
		Exec(context.Background())

	return userHasTicket, err
}
