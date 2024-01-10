package interactors

import (
	"net/http"

	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase/services"
)

type ticketItemInteractor struct {
	db         usecase.DBUsecase
	event      usecase.EventUsecase
	ticketItem usecase.TicketItemUsecase
}

func NewTicketItemInteractor(
	db usecase.DBUsecase,
	event usecase.EventUsecase,
	ticketItem usecase.TicketItemUsecase,
) services.TicketItemService {
	return &ticketItemInteractor{
		db:         db,
		event:      event,
		ticketItem: ticketItem,
	}
}

func (t *ticketItemInteractor) Create(ticket *models.TicketItems) (*models.TicketItemsResponse, *usecase.ResultStatus) {

	db, _ := t.db.Connect()

	_, err := t.event.FindByID(db, ticket.EventID)
	if err != nil {
		return &models.TicketItemsResponse{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	newTicketItem, err := t.ticketItem.Create(db, ticket)
	if err != nil {
		return &models.TicketItemsResponse{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	return newTicketItem.BuildForGet(), usecase.NewResultStatus(http.StatusOK, nil)
}
