package interactors

import (
	"net/http"

	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase/services"
	"github.com/uptrace/bun"
)

type userBookTicketInteractor struct {
	db             usecase.DBUsecase
	event          usecase.EventUsecase
	ticket         usecase.TicketUsecase
	ticketItem     usecase.TicketItemUsecase
	ticketHasItem  usecase.TicketHasItemUsecase
	userBookTicket usecase.UserBookTicketUsecase
	userHasTicket  usecase.UserHasTicketUsecase
}

func NewUserBookTicketInteractor(
	db usecase.DBUsecase,
	event usecase.EventUsecase,
	ticket usecase.TicketUsecase,
	ticketItem usecase.TicketItemUsecase,
	ticketHasItem usecase.TicketHasItemUsecase,
	userBookTicket usecase.UserBookTicketUsecase,
	userHasTicket usecase.UserHasTicketUsecase,
) services.UserBookTicketService {
	return &userBookTicketInteractor{
		db:             db,
		event:          event,
		ticket:         ticket,
		ticketItem:     ticketItem,
		ticketHasItem:  ticketHasItem,
		userBookTicket: userBookTicket,
		userHasTicket:  userHasTicket,
	}
}

func (u *userBookTicketInteractor) Get(bookID string) (*models.UserBookTicketsReponse, *usecase.ResultStatus) {

	db, _ := u.db.Connect()

	userBookTicket, err := u.userBookTicket.FindByBookID(db, bookID)
	if err != nil {
		return &models.UserBookTicketsReponse{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	builtUserBookTicket, err := u.build(db, userBookTicket)
	if err != nil {
		return &models.UserBookTicketsReponse{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	return builtUserBookTicket, usecase.NewResultStatus(http.StatusOK, nil)
}

func (u *userBookTicketInteractor) GetList(userID int, ticketType string) ([]*models.UserBookTicketsReponse, *usecase.ResultStatus) {

	db, _ := u.db.Connect()

	userBookTickets, err := u.userBookTicket.FindByUserID(db, userID)
	if err != nil {
		return []*models.UserBookTicketsReponse{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	builtUserBookTickets := make([]*models.UserBookTicketsReponse, len(userBookTickets))
	for i, userBookTicket := range userBookTickets {
		builtUserBookTicket, err := u.build(db, userBookTicket)
		if err != nil {
			continue
		}

		builtUserBookTickets[i] = builtUserBookTicket
	}

	return builtUserBookTickets, usecase.NewResultStatus(http.StatusOK, nil)
}

func (u *userBookTicketInteractor) build(db bun.IDB, userBookTicket *models.UserBookTickets) (*models.UserBookTicketsReponse, error) {

	event, err := u.event.FindByID(db, userBookTicket.EventID)
	if err != nil {
		return &models.UserBookTicketsReponse{}, err
	}

	userHasTickets, err := u.userHasTicket.FindByUserBookTicketID(db, userBookTicket.ID)
	if err != nil {
		return &models.UserBookTicketsReponse{}, err
	}

	builtUserHasTickets := make([]*models.UserHasTicketsResponse, len(userHasTickets))
	for i, userHasTicket := range userHasTickets {
		builtUserHasTickets[i] = userHasTicket.BuildForGet()
	}

	builtUserBookTicket := userBookTicket.BuildForGet()

	builtUserBookTicket.Event = event.BuildFor()
	builtUserBookTicket.Ticket = u.buildTicket(db, userBookTicket.TicketID, userBookTicket.TicketItemID)
	builtUserBookTicket.UserHasTickets = builtUserHasTickets

	return builtUserBookTicket, nil
}

func (u *userBookTicketInteractor) buildTicket(db bun.IDB, ticketID, ticketItemID int) *models.TicketsResponse {

	ticket, err := u.ticket.FindByID(db, ticketID)
	if err != nil {
		return &models.TicketsResponse{}
	}

	ticketHasItems, _ := u.ticketHasItem.FindByTicketID(db, ticketID)

	ticketItems := make([]*models.TicketItemsResponse, len(ticketHasItems))

	for i, ticketHasItem := range ticketHasItems {
		ticketItem, err := u.ticketItem.FindByID(db, ticketHasItem.TicketItemID)
		if err != nil {
			continue
		}

		ticketItems[i] = ticketItem.BuildForGet()
		ticketItems[i].TicketHasItem = ticketHasItem.BuildForGet()
	}

	builtTicket := ticket.BuildForGet()

	builtTicket.TicketItems = ticketItems

	return builtTicket
}
