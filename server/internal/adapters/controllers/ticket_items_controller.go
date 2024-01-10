package controllers

import (
	"strconv"

	"github.com/takeuchi-shogo/ticket-api/internal/adapters/presenters"
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase/services"
)

type TicketItemsController interface {
	Post(ctx Context)
}

type ticketItemsController struct {
	ticketItem services.TicketItemService
}

func NewTicketItemsController(
	ticketItem services.TicketItemService,
) TicketItemsController {
	return &ticketItemsController{
		ticketItem: ticketItem,
	}
}

func (t *ticketItemsController) Post(ctx Context) {

	eventID, _ := strconv.Atoi(ctx.PostForm("event_id"))
	title := ctx.PostForm("title")
	issuingNumber, _ := strconv.Atoi(ctx.PostForm("issuing_number"))

	ticket, res := t.ticketItem.Create(&models.TicketItems{
		EventID:       eventID,
		Title:         title,
		IssuingNumber: issuingNumber,
	})
	if res.Err != nil {
		ctx.JSON(res.StatusCode, presenters.NewErrResponse(res.Err.Error()))
		return
	}
	ctx.JSON(res.StatusCode, presenters.NewResponse(ticket))
}
