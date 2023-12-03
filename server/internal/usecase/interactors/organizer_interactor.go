package interactors

import (
	"net/http"
	"time"

	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase/services"
	"github.com/takeuchi-shogo/ticket-api/pkg/password"
)

type organizerInteractor struct {
	db        usecase.DBUsecase
	organizer usecase.OrganizerUsecase
}

func NewOrganizerInteractor(db usecase.DBUsecase, organizer usecase.OrganizerUsecase) services.OrganizerService {
	return &organizerInteractor{
		db:        db,
		organizer: organizer,
	}
}

func (o *organizerInteractor) Get(id int) (*models.OrganizersResponse, *usecase.ResultStatus) {
	return &models.OrganizersResponse{}, usecase.NewResultStatus(http.StatusOK, nil)
}

func (o *organizerInteractor) Create(organizer *models.Organizers) (*models.OrganizersResponse, *usecase.ResultStatus) {
	db := o.db.Connect()

	buildPassword, err := password.GenerateFromPassword(organizer.Password)
	if err != nil {
		return &models.OrganizersResponse{}, usecase.NewResultStatus(http.StatusBadRequest, nil)
	}

	organizer.Password = buildPassword
	organizer.CreatedAt = time.Now().Unix()
	organizer.UpdatedAt = time.Now().Unix()

	res, err := o.organizer.Create(db, organizer)
	if err != nil {
		return &models.OrganizersResponse{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}
	return res.BuildFor(), usecase.NewResultStatus(http.StatusOK, nil)
}
