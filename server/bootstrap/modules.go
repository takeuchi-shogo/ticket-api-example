package bootstrap

import (
	"github.com/takeuchi-shogo/ticket-api/config"
	"github.com/takeuchi-shogo/ticket-api/internal/adapters/controllers"
	"github.com/takeuchi-shogo/ticket-api/internal/adapters/gateways"
	"github.com/takeuchi-shogo/ticket-api/internal/adapters/gateways/database"
	"github.com/takeuchi-shogo/ticket-api/internal/infrastructure"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase/interactors"
	"github.com/takeuchi-shogo/ticket-api/pkg/token"
	"go.uber.org/fx"
)

var CommonModule = fx.Options(
	config.Module,
	controllers.Module,
	database.Module,
	gateways.Module,
	interactors.Module,
	infrastructure.Module,
	token.Module,
)
