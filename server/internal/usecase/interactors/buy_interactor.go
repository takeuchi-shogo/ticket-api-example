package interactors

import "github.com/takeuchi-shogo/ticket-api/internal/usecase/services"

type buyInteractor struct{}

func NewBuyInteractor() services.BuyService {
	return &buyInteractor{}
}
