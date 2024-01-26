package controllers

import (
	"strconv"

	"github.com/takeuchi-shogo/ticket-api/internal/adapters/presenters"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase/services"
	"github.com/takeuchi-shogo/ticket-api/pkg/constants"
	"github.com/takeuchi-shogo/ticket-api/pkg/token"
)

type CreditCardsController interface {
	Get(ctx Context)
	Post(ctx Context)
}

type creditCardsController struct {
	creditCard services.CreditCardService
}

func NewCreditCardsController(
	creditCard services.CreditCardService,
) CreditCardsController {
	return &creditCardsController{
		creditCard: creditCard,
	}
}

func (c creditCardsController) Get(ctx Context) {

	authPayload := ctx.MustGet(constants.AuthorizationPayloadKey).(*token.CustomClaims)

	userID, _ := strconv.Atoi(authPayload.UserID)
	creditCard, res := c.creditCard.Get(userID)
	if res.Err != nil {
		ctx.JSON(res.StatusCode, presenters.NewErrResponse(res.Err.Error()))
		return
	}
	ctx.JSON(res.StatusCode, presenters.NewResponse(creditCard))
}

func (c creditCardsController) Post(ctx Context) {

	authPayload := ctx.MustGet(constants.AuthorizationPayloadKey).(*token.CustomClaims)

	userID, _ := strconv.Atoi(authPayload.UserID)

	token := ctx.PostForm("token")
	creditCard, res := c.creditCard.Create(userID, token)
	if res.Err != nil {
		ctx.JSON(res.StatusCode, presenters.NewErrResponse(res.Err.Error()))
		return
	}
	ctx.JSON(res.StatusCode, presenters.NewResponse(creditCard))
}
