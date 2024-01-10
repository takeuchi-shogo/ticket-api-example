package database

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(
		NewAdministratorRepository,
		NewArtistRepository,
		NewEventRepository,
		NewOrganizerRepository,
		NewRegisterEmailRepository,
		NewTicketRepository,
		NewTicketItemRepository,
		NewTicketHasItemRepository,
		NewUserRepository,
	),
)
