package controllers

import "github.com/takeuchi-shogo/ticket-api/internal/usecase/services"

type BuyController interface {
	Post(ctx Context)
}

type buyController struct {
	buy services.BuyService
}

func NewBuyController(
	buy services.BuyService,
) BuyController {
	return &buyController{
		buy: buy,
	}
}

func (b *buyController) Post(ctx Context) {}
