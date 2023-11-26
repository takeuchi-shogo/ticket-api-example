package models

type TicketHasItems struct {
	ID           int     `json:"id"`
	EventID      int     `json:"event_id"`
	TicketID     int     `json:"ticket_id"`
	TicketItemID int     `json:"ticket_item_id"`
	Amount       int64   `json:"amount"`
	Remarks      *string `json:"remarks"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}

type TicketHasItemsResponse struct {
	ID           int     `json:"id"`
	EventID      int     `json:"event_id"`
	TicketID     int     `json:"ticket_id"`
	TicketItemID int     `json:"ticket_item_id"`
	Amount       int64   `json:"amount"`
	Remarks      *string `json:"remarks"`
}
