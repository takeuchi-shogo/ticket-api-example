package controllers

import (
	"github.com/takeuchi-shogo/ticket-api/internal/adapters/gateways/database"
	"github.com/takeuchi-shogo/ticket-api/internal/adapters/presenters"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase/interactors"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase/services"
)

type AdministratorsController interface {
	Get(ctx Context)
}

type administratorsController struct {
	administrator services.AdministratorService
}

func NewAdministratorsController(a services.AdministratorService) AdministratorsController {
	return &administratorsController{
		administrator: &interactors.AdministratorInteractor{
			Administrator: &database.AdministratorRepository{},
		},
	}
}

func (a *administratorsController) Get(ctx Context) {
	ad, res := a.administrator.Get(1)
	if res.Err != nil {
		ctx.JSON(res.StatusCode, presenters.Response{Message: res.Err.Error(), Data: nil})
		return
	}
	ctx.JSON(res.StatusCode, presenters.Response{Message: res.Err.Error(), Data: ad})

}
