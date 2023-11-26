package services

import (
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
)

type AdministratorService interface {
	Get(id int) (*models.Administrators, *usecase.ResultStatus)
	// Create(*models.Administrators) (*models.AdministratorsResponse, *usecase.ResultStatus)
	// Save(*models.Administrators) (*models.AdministratorsResponse, *usecase.ResultStatus)
}
