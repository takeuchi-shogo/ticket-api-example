package services

import (
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
)

type ArtistService interface {
	Get(id int) (*models.ArtistsResponse, *usecase.ResultStatus)
	GetList() ([]*models.ArtistsResponse, *usecase.ResultStatus)
	Create(artist *models.Artists) (*models.ArtistsResponse, *usecase.ResultStatus)
}
