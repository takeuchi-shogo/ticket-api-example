package controllers

import "go.uber.org/fx"

var Module = fx.Options(
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
