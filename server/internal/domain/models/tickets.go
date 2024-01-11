package models

type Tickets struct {
	ID                     int     `json:"id"`
	Title                  string  `json:"title"`
	EventID                int     `json:"event_id"`
	VenueID                *int    `json:"venue_id"`
	PerformanceDate        int64   `json:"performance_date"`
	Note                   *string `json:"note"`
	SaleType               string  `json:"sale_type"`
	StartAt                int64   `json:"start_at"`
	FinishAt               int64   `json:"finish_at"`
	LotteryAt              int64   `json:"lottery_at"`
	IsPaymentByCreditCard  bool    `json:"is_payment_by_credit_card"`
	IsPaymentByConvenience bool    `json:"is_payment_by_convenience"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}

type TicketsResponse struct {
	ID                     int     `json:"id"`
	Title                  string  `json:"title"`
	EventID                int     `json:"event_id"`
	VenueID                *int    `json:"venue_id"`
	PerformanceDate        int64   `json:"performance_date"`
	Note                   *string `json:"note"`
	SaleType               string  `json:"sale_type"`
	StartAt                int64   `json:"start_at"`
	FinishAt               int64   `json:"finish_at"`
	LotteryAt              int64   `json:"lottery_at"`
	IsPaymentByCreditCard  bool    `json:"is_payment_by_credit_card"`
	IsPaymentByConvenience bool    `json:"is_payment_by_convenience"`

	TicketItems []*TicketItemsResponse `json:"ticket_items"`
}

type TicketInteractorResponse struct {
	Tickets []*TicketsResponse `json:"tickets"`
	Total   int                `json:"total"`
}

func (t *Tickets) BuildForGet() *TicketsResponse {
	return &TicketsResponse{
		ID:                     t.ID,
		Title:                  t.Title,
		EventID:                t.EventID,
		VenueID:                t.VenueID,
		PerformanceDate:        t.PerformanceDate,
		Note:                   t.Note,
		SaleType:               t.SaleType,
		StartAt:                t.StartAt,
		FinishAt:               t.FinishAt,
		LotteryAt:              t.LotteryAt,
		IsPaymentByCreditCard:  t.IsPaymentByCreditCard,
		IsPaymentByConvenience: t.IsPaymentByConvenience,
	}
}
