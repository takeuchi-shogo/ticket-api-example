package controllers

import (
	"github.com/takeuchi-shogo/ticket-api/internal/adapters/controllers/tasks"
	"go.uber.org/fx"
)

var Module = fx.Options(
	tasks.Module,
	fx.Provide(
		NewAdministratorsController,
		NewArtistsController,
		NewAuthController,
		NewBuysController,
		NewCreditCardsController,
		NewOrganizersController,
		NewEventsController,
		NewMeController,
		NewTicketsController,
		NewTicketItemsController,
		NewTicketHasItemsController,
		NewUsersController,
	),
)
