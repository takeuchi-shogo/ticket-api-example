package tasks

import (
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase/services/tasks"
)

type paymentInteractor struct {
	db                  usecase.DBUsecase
	paymentByCreditCard usecase.PaymentByCreditCardUsecase
	stripe              usecase.StripeUsecase
	userBookTicket      usecase.UserBookTicketUsecase
}

func NewPaymentInteractor(
	db usecase.DBUsecase,
	paymentByCreditCard usecase.PaymentByCreditCardUsecase,
	stripe usecase.StripeUsecase,
	userBookTicket usecase.UserBookTicketUsecase,
) tasks.PaymentService {
	return &paymentInteractor{
		db:                  db,
		paymentByCreditCard: paymentByCreditCard,
		stripe:              stripe,
		userBookTicket:      userBookTicket,
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
		err := p.stripe.CapturePaymentIntent(payment.PaymentID)
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
