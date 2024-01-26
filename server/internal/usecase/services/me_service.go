package services

import (
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
)

type MeService interface {
	Get(user *models.Users) (*models.MeInteractorResponse, *usecase.ResultStatus)
	GetMe(userID int) (*models.UsersResponse, *usecase.ResultStatus)
	Create(user *models.Users) (*models.MeInteractorResponse, *usecase.ResultStatus)
	Save(user *models.Users) (*models.UsersResponse, *usecase.ResultStatus)
}
