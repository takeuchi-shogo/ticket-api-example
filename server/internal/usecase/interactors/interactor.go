package interactors

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(
		NewAdministratorInteractor,
		NewArtistInteractor,
		NewAuthInteractor,
		NewOrganizerInteractor,
		NewEventInteractor,
		NewMeInteractor,
		NewTicketInteractor,
		NewTicketItemInteractor,
		NewTicketHasItemInteractor,
		NewUserInteractor,
	),
)
