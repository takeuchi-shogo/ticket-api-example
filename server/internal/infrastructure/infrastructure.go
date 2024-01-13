package infrastructure

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(
		NewDB,
		NewRedis,
		NewControllers,
		NewRouting,
		NewStripe,
	),
)
