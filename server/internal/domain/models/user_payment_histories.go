package models

type UserPaymentHistories struct {
	ID               int     `json:"id"`
	EventID          int     `json:"evemt_id"`
	OrganizerID      int     `json:"organizer_id"`
	UserBookTicketID int     `json:"user_book_ticket_id"`
	UserID           int     `json:"user_id"`
	Currency         string  `json:"currency"`
	Amount           float64 `json:"amount"`
	Fee              float64 `json:"fee"`
	ExternalFee      float64 `json:"external_fee"`
	Tax              float64 `json:"tax"`
	ChargeID         string  `json:"charge_id"`
	CreatedAt        int64   `json:"created_at"`
}
