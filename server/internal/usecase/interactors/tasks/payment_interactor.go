package tasks

import (
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase/services/tasks"
)

type paymentInteractor struct {
	db usecase.DBUsecase
}

func NewPaymentInteractor(
	db usecase.DBUsecase,
) tasks.PaymentService {
	return &paymentInteractor{}
}

// クレカ決済を実行する
func (p *paymentInteractor) StartCreditCard() {

	// db, _ := p.db.Connect()
}
