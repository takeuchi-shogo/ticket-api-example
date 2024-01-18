package database

import (
	"context"
	"time"

	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
	"github.com/uptrace/bun"
)

type userBookTicketRepository struct{}

func NewUserBookTicketRepository() usecase.UserBookTicketUsecase {
	return &userBookTicketRepository{}
}

func (u *userBookTicketRepository) FindByID(db bun.IDB, id int) (*models.UserBookTickets, error) {
	userBookTicket := &models.UserBookTickets{}

	if err := db.NewSelect().
		Model(userBookTicket).
		Where("id = ?", id).
		Scan(context.Background()); err != nil {
		return &models.UserBookTickets{}, err
	}
	return userBookTicket, nil
}

func (u *userBookTicketRepository) FindByBookID(db bun.IDB, bookID string) (*models.UserBookTickets, error) {
	userBookTicket := &models.UserBookTickets{}

	if err := db.NewSelect().
		Model(userBookTicket).
		Where("book_id = ?", bookID).
		Scan(context.Background()); err != nil {
		return &models.UserBookTickets{}, err
	}
	return userBookTicket, nil
}

func (u *userBookTicketRepository) FindByUserID(db bun.IDB, userID int) ([]*models.UserBookTickets, error) {
	userBookTickets := []*models.UserBookTickets{}

	if err := db.NewSelect().
		Model(userBookTickets).
		Where("user_id = ?", userID).
		Scan(context.Background()); err != nil {
		return []*models.UserBookTickets{}, err
	}
	return userBookTickets, nil
}

func (u *userBookTicketRepository) FindByTicketID(db bun.IDB, ticketID int) ([]*models.UserBookTickets, error) {
	userBookTickets := []*models.UserBookTickets{}

	if err := db.NewSelect().
		Model(userBookTickets).
		Where("ticket_id = ?", ticketID).
		Scan(context.Background()); err != nil {
		return []*models.UserBookTickets{}, err
	}
	return userBookTickets, nil
}

func (u *userBookTicketRepository) Create(db bun.IDB, userBookTicket *models.UserBookTickets) (*models.UserBookTickets, error) {

	userBookTicket.CreatedAt = time.Now().Unix()
	userBookTicket.UpdatedAt = time.Now().Unix()

	_, err := db.NewInsert().Model(userBookTicket).Exec(context.Background())
	return userBookTicket, err
}
