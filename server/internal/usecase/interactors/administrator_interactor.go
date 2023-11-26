package interactors

import (
	"net/http"

	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase/services"
)

type AdministratorInteractor struct {
	Administrator usecase.AdministratorUsecase
}

func NewAdministratorInteractor(
	Administrator usecase.AdministratorUsecase,
) services.AdministratorService {
	return &AdministratorInteractor{
		Administrator: Administrator,
	}
}

func (a *AdministratorInteractor) Get(id int) (*models.Administrators, *usecase.ResultStatus) {
	adminer, err := a.Administrator.FindByID(id)
	if err != nil {
		return &models.Administrators{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}
	return adminer, usecase.NewResultStatus(http.StatusOK, nil)
}
