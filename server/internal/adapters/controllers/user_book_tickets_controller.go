package controllers

import (
	"strconv"

	"github.com/takeuchi-shogo/ticket-api/internal/adapters/presenters"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase/services"
	"github.com/takeuchi-shogo/ticket-api/pkg/constants"
	"github.com/takeuchi-shogo/ticket-api/pkg/token"
)

type UserBookTicketsController interface {
	Get(ctx Context)
	GetList(ctx Context)
}

type userBookTicketsController struct {
	userBookTicket services.UserBookTicketService
}

func NewUserBookTicketsController(
	userBookTicket services.UserBookTicketService,
) UserBookTicketsController {
	return &userBookTicketsController{
		userBookTicket: userBookTicket,
	}
}

func (u *userBookTicketsController) Get(ctx Context) {

	bookID := ctx.Param("bookId")

	userBookTicket, res := u.userBookTicket.Get(bookID)
	if res.Err != nil {
		ctx.JSON(res.StatusCode, presenters.NewErrResponse(res.Err.Error()))
		return
	}
	ctx.JSON(res.StatusCode, presenters.NewResponse(userBookTicket))
}

func (u *userBookTicketsController) GetList(ctx Context) {
	authPayload := ctx.MustGet(constants.AuthorizationPayloadKey).(*token.CustomClaims)

	userID, _ := strconv.Atoi(authPayload.UserID)

	ticketType := ctx.Query("ticket_type")

	userBookTickets, res := u.userBookTicket.GetList(userID, ticketType)
	if res.Err != nil {
		ctx.JSON(res.StatusCode, presenters.NewErrResponse(res.Err.Error()))
		return
	}
	ctx.JSON(res.StatusCode, presenters.NewResponse(userBookTickets))
}
