package services

import (
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
)

type AuthService interface {
	RegisterEmail(email string) *usecase.ResultStatus
	VerifyCode(*models.RegisterEmails) *usecase.ResultStatus
	Verify(userID int) (*models.Users, *usecase.ResultStatus)
	Login(user *models.Users) (*models.MeInteractorResponse, *usecase.ResultStatus)
	Create(*models.Users) (*models.MeInteractorResponse, *usecase.ResultStatus)
}
