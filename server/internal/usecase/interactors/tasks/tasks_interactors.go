package tasks

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(
		NewDrawInteractor,
		NewPaymentInteractor,
		NewMailInteractor,
	),
)
