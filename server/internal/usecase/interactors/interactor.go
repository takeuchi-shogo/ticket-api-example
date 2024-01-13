package interactors

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(
		NewAdministratorInteractor,
		NewArtistInteractor,
		NewAuthInteractor,
		NewBuyInteractor,
		NewCreditCardInteractor,
		NewOrganizerInteractor,
		NewEventInteractor,
		NewMeInteractor,
		NewTicketInteractor,
		NewTicketItemInteractor,
		NewTicketHasItemInteractor,
		NewUserInteractor,
	),
)
