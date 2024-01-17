package tasks

type PaymentsController interface {
	Start()
}

type paymentsController struct{}

func NewPaymentsController() PaymentsController {
	return &paymentsController{}
}

func (p *paymentsController) Start() {}
