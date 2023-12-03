package controllers

import (
	"net/http"

	"github.com/takeuchi-shogo/ticket-api/internal/adapters/presenters"
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase/services"
)

type OrganizersController interface {
	Get(ctx Context)
	Post(ctx Context)
}

type organizersController struct {
	organizer services.OrganizerService
}

func NewOrganizersController(organizer services.OrganizerService) OrganizersController {
	return &organizersController{
		organizer: organizer,
	}
}

func (o *organizersController) Get(ctx Context) {
	organizer, res := o.organizer.Get(1)
	if res.Err != nil {
		return
	}
	ctx.JSON(res.StatusCode, presenters.Response{Message: "success", Data: organizer})
}

func (o *organizersController) Post(ctx Context) {

	organizer := &models.Organizers{}
	if err := ctx.ShouldBind(organizer); err != nil {
		ctx.JSON(http.StatusBadRequest, presenters.NewErrResponse(err.Error()))
		return
	}

	newOrganizer, res := o.organizer.Create(organizer)
	if res.Err != nil {
		ctx.JSON(res.StatusCode, presenters.NewErrResponse(res.Err.Error()))
		return
	}

	ctx.JSON(res.StatusCode, presenters.NewResponse(newOrganizer))
}
