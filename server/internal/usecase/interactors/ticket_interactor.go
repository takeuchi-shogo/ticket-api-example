package interactors

import (
	"net/http"

	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase/services"
)

type TicketInteractor struct {
	ticket usecase.TicketUsecase
}

func NewTicketInteractor(ticket usecase.TicketUsecase) services.TicketService {
	return &TicketInteractor{
		ticket: ticket,
	}
}

func (t *TicketInteractor) Get(id int) (*models.TicketsResponse, *usecase.ResultStatus) {
	return &models.TicketsResponse{}, usecase.NewResultStatus(http.StatusOK, nil)
}
