package interactors

import (
	"github.com/takeuchi-shogo/ticket-api/internal/usecase/interactors/tasks"
	"go.uber.org/fx"
)

var Module = fx.Options(
	tasks.Module,
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
		NewUserBookTicketInteractor,
	),
)
