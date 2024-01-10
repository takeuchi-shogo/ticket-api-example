package controllers

import (
	"strconv"

	"github.com/takeuchi-shogo/ticket-api/internal/adapters/presenters"
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase/services"
)

type TicketsController interface {
	Get(ctx Context)
	Post(ctx Context)
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
	id, _ := strconv.Atoi(ctx.Param("id"))

	ticket, res := t.ticket.Get(id)
	if res.Err != nil {
		ctx.JSON(res.StatusCode, presenters.ErrResponse{ErrorMessage: res.Err.Error()})
		return
	}
	ctx.JSON(res.StatusCode, presenters.Response{Message: "success", Data: ticket})
}

func (t *ticketsController) Post(ctx Context) {

	title := ctx.PostForm("title")
	eventID, _ := strconv.Atoi(ctx.PostForm("event_id"))
	venueID, _ := strconv.Atoi("venue_id")
	note := ctx.PostForm("note")
	saleType := ctx.PostForm("sale_type")
	startAt, _ := strconv.ParseInt(ctx.PostForm("start_at"), 10, 64)
	finishAt, _ := strconv.ParseInt(ctx.PostForm("finish_at"), 10, 64)
	lotteryAt, _ := strconv.ParseInt(ctx.PostForm("lottery_at"), 10, 64)
	isPaymentByCreditCard, _ := strconv.ParseBool(ctx.PostForm("is_payment_by_credit_card"))
	isPaymentByConvenience, _ := strconv.ParseBool(ctx.PostForm("is_payment_by_convenience"))

	ticket, res := t.ticket.Create(&models.Tickets{
		Title:                  title,
		EventID:                eventID,
		VenueID:                &venueID,
		Note:                   &note,
		SaleType:               saleType,
		StartAt:                startAt,
		FinishAt:               finishAt,
		LotteryAt:              lotteryAt,
		IsPaymentByCreditCard:  isPaymentByCreditCard,
		IsPaymentByConvenience: isPaymentByConvenience,
	})
	if res.Err != nil {
		ctx.JSON(res.StatusCode, presenters.ErrResponse{ErrorMessage: res.Err.Error()})
		return
	}
	ctx.JSON(res.StatusCode, presenters.Response{Message: "success", Data: ticket})
}
