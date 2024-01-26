package models

type TicketItems struct {
	ID            int    `json:"id"`
	EventID       int    `json:"event_id"`
	Title         string `json:"title"`
	IssuingNumber int    `json:"issuing_number"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}

type TicketItemsResponse struct {
	ID            int    `json:"id"`
	EventID       int    `json:"event_id"`
	Title         string `json:"title"`
	IssuingNumber int    `json:"issuing_number"`

	// TicketHasItem 情報も加えておく
	TicketHasItem *TicketHasItemsResponse `json:"ticket_has_item"`
}

func (t *TicketItems) BuildForGet() *TicketItemsResponse {
	return &TicketItemsResponse{
		ID:            t.ID,
		EventID:       t.EventID,
		Title:         t.Title,
		IssuingNumber: t.IssuingNumber,
	}
}
