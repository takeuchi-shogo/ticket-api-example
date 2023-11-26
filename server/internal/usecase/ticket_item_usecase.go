package usecase

import "github.com/takeuchi-shogo/ticket-api/internal/domain/models"

type TicketItemUsecase interface {
	FindByID(id int) (*models.TicketItems, error)
}
