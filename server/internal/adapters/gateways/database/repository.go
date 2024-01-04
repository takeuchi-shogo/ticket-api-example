package database

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(
		NewAdministratorRepository,
		NewEventRepository,
		NewOrganizerRepository,
		NewRegisterEmailRepository,
		NewTicketRepository,
		NewUserRepository,
	),
)
