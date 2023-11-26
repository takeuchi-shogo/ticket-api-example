package controllers

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(
		NewAdministratorsController,
		NewAuthController,
		NewOrganizersController,
		NewEventsController,
		NewTicketsController,
		NewUsersController,
	),
)
