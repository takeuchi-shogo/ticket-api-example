package usecase

import "github.com/takeuchi-shogo/ticket-api/internal/domain/models"

type UserUsecase interface {
	FindByID(id int) (*models.Users, error)
}
