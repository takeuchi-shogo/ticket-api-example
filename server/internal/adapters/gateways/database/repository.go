package database

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(
		NewAdministratorRepository,
		NewArtistRepository,
		NewCreditCardRepository,
		NewEventRepository,
		NewOrganizerRepository,
		NewRegisterEmailRepository,
		NewPaymentByCreditCardRepository,
		NewTicketRepository,
		NewTicketItemRepository,
		NewTicketHasItemRepository,
		NewUserBookTicketRepository,
		NewUserHasTicketRepository,
		NewUserRepository,
		NewUserMailLogRepository,
	),
)
