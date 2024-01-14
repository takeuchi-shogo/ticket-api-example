package interactors

import (
	"net/http"

	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase/services"
)

type userBookTicketInteractor struct {
	db             usecase.DBUsecase
	userBookTicket usecase.UserBookTicketUsecase
}

func NewUserBookTicketInteractor(
	db usecase.DBUsecase,
	userBookTicket usecase.UserBookTicketUsecase,
) services.UserBookTicketService {
	return &userBookTicketInteractor{
		db:             db,
		userBookTicket: userBookTicket,
	}
}

func (u *userBookTicketInteractor) Get(bookID string) (*models.UserBookTicketsReponse, *usecase.ResultStatus) {

	db, _ := u.db.Connect()

	userBookTicket, err := u.userBookTicket.FindByBookID(db, bookID)
	if err != nil {
		return &models.UserBookTicketsReponse{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	return userBookTicket.BuildForGet(), usecase.NewResultStatus(http.StatusOK, nil)
}

func (u *userBookTicketInteractor) GetList(userID int) ([]*models.UserBookTicketsReponse, *usecase.ResultStatus) {

	db, _ := u.db.Connect()

	userBookTickets, err := u.userBookTicket.FindByUserID(db, userID)
	if err != nil {
		return []*models.UserBookTicketsReponse{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	builtUserBookTickets := make([]*models.UserBookTicketsReponse, len(userBookTickets))
	for i, userBookTicket := range userBookTickets {
		userBookTickets[i] = (*models.UserBookTickets)(userBookTicket.BuildForGet())
	}

	return builtUserBookTickets, usecase.NewResultStatus(http.StatusOK, nil)
}
