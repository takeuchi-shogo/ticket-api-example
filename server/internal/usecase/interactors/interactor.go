package interactors

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(
		NewAdministratorInteractor,
		NewAuthInteractor,
		NewOrganizerInteractor,
		NewEventInteractor,
		NewMeInteractor,
		NewTicketInteractor,
		NewUserInteractor,
	),
)
