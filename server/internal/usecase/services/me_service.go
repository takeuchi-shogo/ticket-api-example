package services

import (
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
)

type MeService interface {
	Get(user *models.Users) (*models.MeInteractorResponse, *usecase.ResultStatus)
	Create(user *models.Users) (*models.MeInteractorResponse, *usecase.ResultStatus)
}
