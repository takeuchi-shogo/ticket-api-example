package models

type UserHasTickets struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	UserBookTicketID int    `json:"user_book_ticket_id"`
	SeatID           *int   `json:"seat_id"`
	TicketStatus     string `json:"ticket_status"`
	ReferenceNumber  *int   `json:"reference_number"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}
