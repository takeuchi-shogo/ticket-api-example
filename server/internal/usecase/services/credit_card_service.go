package services

import (
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
)

type CreditCardService interface {
	Get(userID int) (*models.CreditCardsReponse, *usecase.ResultStatus)
	Create(userID int, token string) (*models.CreditCardsReponse, *usecase.ResultStatus)
}
