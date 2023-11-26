package controllers

import (
	"github.com/takeuchi-shogo/ticket-api/internal/adapters/presenters"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase/services"
)

type TicketsController interface {
	Get(ctx Context)
}

type ticketsController struct {
	ticket services.TicketService
}

func NewTicketsController(ticket services.TicketService) TicketsController {
	return &ticketsController{
		ticket: ticket,
	}
}

func (t *ticketsController) Get(ctx Context) {
	ticket, res := t.ticket.Get(1)
	if res.Err != nil {
		ctx.JSON(res.StatusCode, presenters.ErrResponse{ErrorMessage: res.Err.Error()})
		return
	}
	ctx.JSON(res.StatusCode, presenters.Response{Message: "success", Data: ticket})
}
