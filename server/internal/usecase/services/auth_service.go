package services

import (
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
)

type AuthService interface {
	Verify(userID int) (*models.Users, *usecase.ResultStatus)
	Create(*models.Users) (*models.Users, *usecase.ResultStatus)
}
