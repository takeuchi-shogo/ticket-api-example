package interactors

import (
	"net/http"

	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase/services"
	"github.com/uptrace/bun"
)

type ticketHasItemInteractor struct {
	db            usecase.DBUsecase
	event         usecase.EventUsecase
	ticket        usecase.TicketUsecase
	ticketItem    usecase.TicketItemUsecase
	ticketHasItem usecase.TicketHasItemUsecase
}

func NewTicketHasItemInteractor(
	db usecase.DBUsecase,
	event usecase.EventUsecase,
	ticket usecase.TicketUsecase,
	ticketItem usecase.TicketItemUsecase,
	ticketHasItem usecase.TicketHasItemUsecase,
) services.TicketHasItemService {
	return &ticketHasItemInteractor{
		db:            db,
		event:         event,
		ticket:        ticket,
		ticketItem:    ticketItem,
		ticketHasItem: ticketHasItem,
	}
}

func (t *ticketHasItemInteractor) Create(ticketHasItem *models.TicketHasItems) (*models.TicketHasItemsResponse, *usecase.ResultStatus) {

	db, _ := t.db.Connect()

	if err := t.canCreate(db, ticketHasItem); err != nil {
		return &models.TicketHasItemsResponse{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	newTicketHasItem, err := t.ticketHasItem.Create(db, ticketHasItem)
	if err != nil {
		return &models.TicketHasItemsResponse{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}
	return newTicketHasItem.BuildForGet(), usecase.NewResultStatus(http.StatusOK, nil)
}

func (t *ticketHasItemInteractor) canCreate(db bun.IDB, ticketHasItem *models.TicketHasItems) error {

	_, err := t.event.FindByID(db, ticketHasItem.EventID)
	if err != nil {
		return err
	}

	_, err = t.ticket.FindByID(db, ticketHasItem.TicketID)
	if err != nil {
		return err
	}
	_, err = t.ticketItem.FindByID(db, ticketHasItem.TicketItemID)
	if err != nil {
		return err
	}

	return nil
}
