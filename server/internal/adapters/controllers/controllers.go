package controllers

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(
		NewAdministratorsController,
		NewArtistsController,
		NewAuthController,
		NewOrganizersController,
		NewEventsController,
		NewMeController,
		NewTicketsController,
		NewTicketItemsController,
		NewTicketHasItemsController,
		NewUsersController,
	),
)
