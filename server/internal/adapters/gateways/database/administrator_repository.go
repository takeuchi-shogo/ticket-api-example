package database

import (
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
)

type AdministratorRepository struct{}

func NewAdministratorRepository() usecase.AdministratorUsecase {
	return &AdministratorRepository{}
}

func (a *AdministratorRepository) FindByID(id int) (*models.Administrators, error) {
	return &models.Administrators{}, nil
}
