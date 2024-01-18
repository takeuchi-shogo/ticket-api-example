package models

type UserBookTickets struct {
	ID              int    `json:"id" bun:",pk,autoincrement"`
	BookID          string `json:"book_id"`
	UserID          int    `json:"user_id"`
	EventID         int    `json:"event_id"`
	TicketID        int    `json:"ticket_id"`
	TicketItemID    int    `json:"ticket_item_id"`
	PaymentMethod   string `json:"payment_method"`
	NumberOfTickets int    `json:"number_of_tickets"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}

type UserBookTicketsReponse struct {
	ID              int    `json:"id"`
	BookID          string `json:"book_id"`
	UserID          int    `json:"user_id"`
	EventID         int    `json:"event_id"`
	TicketID        int    `json:"ticket_id"`
	TicketItemID    int    `json:"ticket_item_id"`
	PaymentMethod   string `json:"payment_method"`
	NumberOfTickets int    `json:"number_of_tickets"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}

func (u *UserBookTickets) BuildForGet() *UserBookTicketsReponse {
	return &UserBookTicketsReponse{
		ID:              u.ID,
		BookID:          u.BookID,
		UserID:          u.UserID,
		EventID:         u.EventID,
		TicketID:        u.TicketID,
		TicketItemID:    u.TicketItemID,
		PaymentMethod:   u.PaymentMethod,
		NumberOfTickets: u.NumberOfTickets,
		CreatedAt:       u.CreatedAt,
		UpdatedAt:       u.UpdatedAt,
	}
}
