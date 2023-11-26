package usecase

import "github.com/takeuchi-shogo/ticket-api/internal/domain/models"

type AdministratorUsecase interface {
	FindByID(id int) (*models.Administrators, error)
}
