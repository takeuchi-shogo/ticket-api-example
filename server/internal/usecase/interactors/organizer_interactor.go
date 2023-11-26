package interactors

import (
	"net/http"

	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase/services"
)

type organizerInteractor struct {
	organizer usecase.OrganizerUsecase
}

func NewOrganizerInteractor(organizer usecase.OrganizerUsecase) services.OrganizerService {
	return &organizerInteractor{
		organizer: organizer,
	}
}

func (o *organizerInteractor) Get(id int) (*models.OrganizersResponse, *usecase.ResultStatus) {
	return &models.OrganizersResponse{}, usecase.NewResultStatus(http.StatusOK, nil)
}

func (o *organizerInteractor) Create(*models.Organizers) (*models.OrganizersResponse, *usecase.ResultStatus) {
	return &models.OrganizersResponse{}, usecase.NewResultStatus(http.StatusOK, nil)
}
