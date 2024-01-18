package interactors

import (
	"net/http"

	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase/services"
	"github.com/takeuchi-shogo/ticket-api/pkg/random"
	"github.com/uptrace/bun"
)

type buyInteractor struct {
	creditCard          usecase.CreditCardUsecase
	db                  usecase.DBUsecase
	paymentByCreditCard usecase.PaymentByCreditCardUsecase
	stripe              usecase.StripeUsecase
	ticket              usecase.TicketUsecase
	ticketItem          usecase.TicketItemUsecase
	ticketHasItem       usecase.TicketHasItemUsecase
	userBookTicket      usecase.UserBookTicketUsecase
	userHasTicket       usecase.UserHasTicketUsecase
}

func NewBuyInteractor(
	creditCard usecase.CreditCardUsecase,
	db usecase.DBUsecase,
	paymentByCreditCard usecase.PaymentByCreditCardUsecase,
	stripe usecase.StripeUsecase,
	ticket usecase.TicketUsecase,
	ticketItem usecase.TicketItemUsecase,
	ticketHasItem usecase.TicketHasItemUsecase,
	userBookTicket usecase.UserBookTicketUsecase,
	userHasTicket usecase.UserHasTicketUsecase,
) services.BuyService {
	return &buyInteractor{
		creditCard:          creditCard,
		db:                  db,
		paymentByCreditCard: paymentByCreditCard,
		stripe:              stripe,
		ticket:              ticket,
		ticketItem:          ticketItem,
		ticketHasItem:       ticketHasItem,
		userBookTicket:      userBookTicket,
		userHasTicket:       userHasTicket,
	}
}

func (b *buyInteractor) Create(userBookTicket *models.UserBookTickets) (*models.UserBookTicketsReponse, *usecase.ResultStatus) {

	db, _ := b.db.Transaction()

	// 入力内容の確認
	amount, err := b.validateParameterAndCalculateAmount(db, userBookTicket)
	if err != nil {
		db.Rollback()
		return &models.UserBookTicketsReponse{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}
	// UserBookTicketの作成

	userBookTicket.BookID = b.randomBookID(db)
	newUserBookTicket, err := b.userBookTicket.Create(db, userBookTicket)
	if err != nil {
		db.Rollback()
		return &models.UserBookTicketsReponse{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	// UserHasTicketの作成
	userHasTickets := make([]*models.UserHasTickets, userBookTicket.NumberOfTickets)
	for i := 0; i < newUserBookTicket.NumberOfTickets; i++ {
		userHasTicket, err := b.userHasTicket.Create(db, &models.UserHasTickets{
			UserID:           newUserBookTicket.UserID,
			UserBookTicketID: newUserBookTicket.ID,
			SeatID:           nil,
			TicketStatus:     "lottery",
			ReferenceNumber:  nil,
		})
		if err != nil {
			db.Rollback()
			return &models.UserBookTicketsReponse{}, usecase.NewResultStatus(http.StatusBadRequest, err)
		}
		userHasTickets[i] = userHasTicket
	}

	// Payment
	// Stripeへの決済
	_, err = b.paymentToStripeByCreditCard(db, amount, newUserBookTicket.UserID, newUserBookTicket.ID)
	if err != nil {
		db.Rollback()
		return &models.UserBookTicketsReponse{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	db.Commit()
	return newUserBookTicket.BuildForGet(), usecase.NewResultStatus(http.StatusCreated, err)
}

func (b *buyInteractor) validateParameterAndCalculateAmount(db bun.IDB, userBookTicket *models.UserBookTickets) (int, error) {

	// 販売スケジュールの確認
	ticket, err := b.ticket.FindByID(db, userBookTicket.TicketID)
	if err != nil {
		return 0, err
	}

	// チケットの確認
	ticketItem, err := b.ticketItem.FindByID(db, userBookTicket.TicketItemID)
	if err != nil {
		return 0, err
	}

	// スケジュールに紐付いたチケット詳細情報の取得
	ticketHasItem, err := b.ticketHasItem.FindByTicketIDAndTicketItemID(db, ticket.ID, ticketItem.ID)
	if err != nil {
		return 0, err
	}

	// チケット料金の計算
	amount := ticketHasItem.Amount * float64(userBookTicket.NumberOfTickets)
	return int(amount), nil
}

func (b *buyInteractor) paymentToStripeByCreditCard(db bun.IDB, amount int, userID, userBookTicketID int) (*models.PaymentByCreditCards, error) {

	creditCard, err := b.creditCard.FindByUserID(db, userID)
	if err != nil {
		return &models.PaymentByCreditCards{}, err
	}
	// カスタマー情報の取得
	// customer, err := b.stripe.GetCustomer(creditCard.CustomerID)
	// if err != nil {
	// 	return &models.PaymentByCreditCards{}, err
	// }

	// オーソリ（与信の確保）
	paymentID, err := b.stripe.AuthenticatePaymentIntent(amount, creditCard.CustomerID)
	if err != nil {
		return &models.PaymentByCreditCards{}, err
	}

	paymentBycreditCard, err := b.paymentByCreditCard.Create(db, &models.PaymentByCreditCards{
		UserBookTicketID: userBookTicketID,
		PaymentID:        paymentID,
	})
	if err != nil {
		return &models.PaymentByCreditCards{}, err
	}

	return paymentBycreditCard, nil
}

func (b *buyInteractor) randomBookID(db bun.IDB) string {

	bookID := random.RandomString(10)
	for {
		_, err := b.userBookTicket.FindByBookID(db, bookID)
		if err != nil {
			break
		}
		bookID = random.RandomString(10)
	}

	return bookID
}
