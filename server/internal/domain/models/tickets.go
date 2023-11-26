package models

type Tickets struct {
	ID       int     `json:"id"`
	Title    string  `json:"title"`
	Note     *string `json:"note"`
	SaleType string  `json:"sale_type"`
	StartAt  int64   `json:"start_at"`
	FinishAt int64   `json:"finish_at"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}

type TicketsResponse struct {
	ID       int     `json:"id"`
	Title    string  `json:"title"`
	Note     *string `json:"note"`
	SaleType string  `json:"sale_type"`
	StartAt  int64   `json:"start_at"`
	FinishAt int64   `json:"finish_at"`
}
