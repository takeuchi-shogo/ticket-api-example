package tasks

import (
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase/services/tasks"
	"github.com/uptrace/bun"
)

type paymentInteractor struct {
	db                  usecase.DBUsecase
	event               usecase.EventUsecase
	paymentByCreditCard usecase.PaymentByCreditCardUsecase
	stripe              usecase.StripeUsecase
	userBookTicket      usecase.UserBookTicketUsecase
	userPaymentHistory  usecase.UserPaymentHistoryUsecase
}

func NewPaymentInteractor(
	db usecase.DBUsecase,
	event usecase.EventUsecase,
	paymentByCreditCard usecase.PaymentByCreditCardUsecase,
	stripe usecase.StripeUsecase,
	userBookTicket usecase.UserBookTicketUsecase,
	userPaymentHistory usecase.UserPaymentHistoryUsecase,
) tasks.PaymentService {
	return &paymentInteractor{
		db:                  db,
		event:               event,
		paymentByCreditCard: paymentByCreditCard,
		stripe:              stripe,
		userBookTicket:      userBookTicket,
		userPaymentHistory:  userPaymentHistory,
	}
}

// クレカ決済を実行する
func (p *paymentInteractor) StartCreditCard() {

	db, _ := p.db.Connect()

	// まだ支払いの済んでいないものを取得
	payments, _ := p.paymentByCreditCard.FindByIsValid(db, true)
	if len(payments) <= 0 {
		return
	}

	for _, payment := range payments {
		// Stripeへ決済する
		result, err := p.stripe.CapturePaymentIntent(payment.PaymentID)
		if err != nil {
			continue
		}

		// 購入履歴を作成する
		err = p.createUserPaymentHistory(db, payment, result)
		if err != nil {
			continue
		}

		// 完了した決済は無効にしておく
		payment.IsValid = false
		_, err = p.paymentByCreditCard.Save(db, payment)
		if err != nil {
			continue
		}
	}
}

func (p *paymentInteractor) createUserPaymentHistory(db bun.IDB, payment *models.PaymentByCreditCards, result *models.StripePaymentIntents) error {

	// UserBookTicketの取得
	userBookTicket, err := p.userBookTicket.FindByID(db, payment.UserBookTicketID)
	if err != nil {
		return err
	}

	event, err := p.event.FindByID(db, userBookTicket.EventID)
	if err != nil {
		return err
	}

	// UserPaymentHistoryの作成
	_, err = p.userPaymentHistory.Create(db, &models.UserPaymentHistories{
		EventID:          userBookTicket.EventID,
		OrganizerID:      event.OrganizerID,
		UserBookTicketID: payment.UserBookTicketID,
		UserID:           userBookTicket.UserID,
		Amount:           float64(result.Amount),
		ChargeID:         result.ID,
	})
	if err != nil {
		return err
	}
	return nil
}
