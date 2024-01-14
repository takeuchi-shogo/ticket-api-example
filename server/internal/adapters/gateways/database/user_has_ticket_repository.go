package database

import (
	"context"

	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
	"github.com/uptrace/bun"
)

type userHasTicketRepository struct{}

func NewUserHasTicketRepository() usecase.UserHasTicketUsecase {
	return &userHasTicketRepository{}
}

func (u *userHasTicketRepository) Create(db bun.IDB, userHasTicket *models.UserHasTickets) (*models.UserHasTickets, error) {
	_, err := db.NewInsert().Model(userHasTicket).Exec(context.Background())
	return userHasTicket, err
}
