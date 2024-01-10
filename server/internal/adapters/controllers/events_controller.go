package controllers

import (
	"net/http"
	"strconv"

	"github.com/takeuchi-shogo/ticket-api/internal/adapters/presenters"
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase/services"
)

type EventController interface {
	Get(ctx Context)
	GetList(ctx Context)
	GetListByArtistID(ctx Context)
	Post(ctx Context)
}

type eventsController struct {
	eventService services.EventService
}

func NewEventsController(event services.EventService) EventController {
	return &eventsController{
		eventService: event,
	}
}

func (e *eventsController) Get(ctx Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	event, res := e.eventService.Get(id)
	if res.Err != nil {
		ctx.JSON(res.StatusCode, presenters.Response{Message: res.Err.Error(), Data: nil})
		return
	}

	ctx.JSON(res.StatusCode, presenters.Response{Message: "success", Data: event})
}

func (e *eventsController) GetList(ctx Context) {

	events, res := e.eventService.GetList()
	if res.Err != nil {
		ctx.JSON(res.StatusCode, presenters.Response{Message: res.Err.Error(), Data: nil})
		return
	}

	ctx.JSON(res.StatusCode, presenters.Response{Message: "success", Data: events})
}

func (e *eventsController) GetListByArtistID(ctx Context) {

	eventID, _ := strconv.Atoi(ctx.Param("eventID"))
	artistID, _ := strconv.Atoi(ctx.Param("artistID"))

	events, res := e.eventService.GetListByArtistID(eventID, artistID)
	if res.Err != nil {
		ctx.JSON(res.StatusCode, presenters.Response{Message: res.Err.Error(), Data: nil})
		return
	}

	ctx.JSON(res.StatusCode, presenters.Response{Message: "success", Data: events})
}

func (e *eventsController) Post(ctx Context) {

	event := &models.Events{}
	if err := ctx.BindJSON(event); err != nil {
		ctx.JSON(http.StatusBadRequest, presenters.Response{Message: err.Error(), Data: nil})
		return
	}

	newEvent, res := e.eventService.Create(event)
	if res.Err != nil {
		ctx.JSON(res.StatusCode, presenters.Response{Message: res.Err.Error(), Data: nil})
		return
	}

	ctx.JSON(res.StatusCode, presenters.Response{Message: "success", Data: newEvent})
}
