package interactors

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(
		NewAuthInteractor,
		NewEventInteractor,
		NewUserInteractor,
	),
)
