package controllers

import (
	"strconv"

	"github.com/takeuchi-shogo/ticket-api/internal/adapters/presenters"
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase/services"
)

type TicketHasItemsController interface {
	Post(ctx Context)
}

type ticketHasItemsController struct {
	ticketHasItem services.TicketHasItemService
}

func NewTicketHasItemsController(ticketHasItem services.TicketHasItemService) TicketHasItemsController {
	return &ticketHasItemsController{
		ticketHasItem: ticketHasItem,
	}
}

func (t *ticketHasItemsController) Post(ctx Context) {

	eventID, _ := strconv.Atoi(ctx.PostForm("event_id"))
	ticketID, _ := strconv.Atoi(ctx.PostForm("ticket_id"))
	ticketItemID, _ := strconv.Atoi(ctx.PostForm("ticket_item_id"))
	amount, _ := strconv.ParseInt(ctx.PostForm("amount"), 10, 64)
	remarks := ctx.PostForm("remarks")

	ticket, res := t.ticketHasItem.Create(&models.TicketHasItems{
		EventID:      eventID,
		TicketID:     ticketID,
		TicketItemID: ticketItemID,
		Amount:       float64(amount),
		Remarks:      &remarks,
	})
	if res.Err != nil {
		ctx.JSON(res.StatusCode, presenters.ErrResponse{ErrorMessage: res.Err.Error()})
		return
	}
	ctx.JSON(res.StatusCode, presenters.Response{Message: "success", Data: ticket})
}
