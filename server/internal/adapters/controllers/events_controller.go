package controllers

import (
	"github.com/takeuchi-shogo/ticket-api/internal/adapters/presenters"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase/services"
)

type EventController interface {
	Get(ctx Context)
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
	event, res := e.eventService.Get(1)
	if res.Err != nil {
		ctx.JSON(res.StatusCode, presenters.Response{Message: res.Err.Error(), Data: nil})
		return
	}
	ctx.JSON(res.StatusCode, presenters.Response{Message: "success", Data: event})

}
