package controllers

import (
	"strconv"

	"github.com/takeuchi-shogo/ticket-api/internal/adapters/presenters"
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase/services"
	"github.com/takeuchi-shogo/ticket-api/pkg/constants"
	"github.com/takeuchi-shogo/ticket-api/pkg/token"
)

type BuysController interface {
	Post(ctx Context)
}

type buysController struct {
	buy services.BuyService
}

func NewBuysController(
	buy services.BuyService,
) BuysController {
	return &buysController{
		buy: buy,
	}
}

func (b *buysController) Post(ctx Context) {

	authPayload := ctx.MustGet(constants.AuthorizationPayloadKey).(*token.CustomClaims)

	userID, _ := strconv.Atoi(authPayload.UserID)

	eventID, _ := strconv.Atoi(ctx.PostForm("event_id"))
	ticketID, _ := strconv.Atoi(ctx.PostForm("ticket_id"))
	ticketItemID, _ := strconv.Atoi(ctx.PostForm("ticket_item_id"))
	paymentMethod := ctx.PostForm("payment_method")
	numberOfTickets, _ := strconv.Atoi(ctx.PostForm("number_of_tickets"))

	userBookTicket, res := b.buy.Create(&models.UserBookTickets{
		UserID:          userID,
		EventID:         eventID,
		TicketID:        ticketID,
		TicketItemID:    ticketItemID,
		PaymentMethod:   paymentMethod,
		NumberOfTickets: numberOfTickets,
	})
	if res.Err != nil {
		ctx.JSON(res.StatusCode, presenters.NewErrResponse(res.Err.Error()))
		return
	}
	ctx.JSON(res.StatusCode, presenters.NewResponse(userBookTicket))
}
