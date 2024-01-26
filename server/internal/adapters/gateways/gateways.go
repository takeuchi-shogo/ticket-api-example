package gateways

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(
		NewDBGateway,
		NewStripeGateway,
	),
)
