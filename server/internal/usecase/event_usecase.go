package usecase

import "github.com/takeuchi-shogo/ticket-api/internal/domain/models"

type EventUsecase interface {
	FindByID(id int) (*models.Events, error)
}
