package interactors

import (
	"net/http"

	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase/services"
	"github.com/uptrace/bun"
)

type ticketInteractor struct {
	db            usecase.DBUsecase
	ticket        usecase.TicketUsecase
	ticketItem    usecase.TicketItemUsecase
	ticketHasItem usecase.TicketHasItemUsecase
}

func NewTicketInteractor(
	db usecase.DBUsecase,
	ticket usecase.TicketUsecase,
	ticketItem usecase.TicketItemUsecase,
	ticketHasItem usecase.TicketHasItemUsecase,
) services.TicketService {
	return &ticketInteractor{
		db:            db,
		ticket:        ticket,
		ticketItem:    ticketItem,
		ticketHasItem: ticketHasItem,
	}
}

func (t *ticketInteractor) Get(id int) (*models.TicketsResponse, *usecase.ResultStatus) {

	db, _ := t.db.Connect()

	ticket, err := t.ticket.FindByID(db, id)
	if err != nil {
		return &models.TicketsResponse{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	builtTicket, err := t.build(db, ticket)
	if err != nil {
		return &models.TicketsResponse{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	return builtTicket, usecase.NewResultStatus(http.StatusOK, nil)
}

func (t *ticketInteractor) build(db bun.IDB, ticket *models.Tickets) (*models.TicketsResponse, error) {

	builtTicket := ticket.BuildForGet()

	ticketHasItems, _ := t.ticketHasItem.FindByTicketID(db, ticket.ID)

	ticketItems := make([]*models.TicketItemsResponse, len(ticketHasItems))

	for i, ticketHasItem := range ticketHasItems {
		ticketItem, err := t.ticketItem.FindByID(db, ticketHasItem.TicketItemID)
		if err != nil {
			continue
		}

		ticketItems[i] = ticketItem.BuildForGet()
		ticketItems[i].TicketHasItem = ticketHasItem.BuildForGet()
	}

	builtTicket.TicketItems = ticketItems

	return builtTicket, nil
}

func (t *ticketInteractor) Create(ticket *models.Tickets) (*models.TicketsResponse, *usecase.ResultStatus) {

	db, _ := t.db.Connect()

	newTicket, err := t.ticket.Create(db, ticket)
	if err != nil {
		return &models.TicketsResponse{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	return newTicket.BuildForGet(), usecase.NewResultStatus(http.StatusCreated, nil)
}
