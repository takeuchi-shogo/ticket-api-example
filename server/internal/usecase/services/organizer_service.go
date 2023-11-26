package services

import (
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
)

type OrganizerService interface {
	Get(id int) (*models.OrganizersResponse, *usecase.ResultStatus)
	Create(*models.Organizers) (*models.OrganizersResponse, *usecase.ResultStatus)
}
