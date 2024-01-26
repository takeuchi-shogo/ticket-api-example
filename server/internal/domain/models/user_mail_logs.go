package models

type UserMailLogs struct {
	ID               int     `json:"id"`
	Email            *string `json:"email"`
	UserID           *int    `json:"user_id"`
	UserBookTicketID *int    `json:"user_book_ticket_id"`
	MailType         string  `json:"mail_type"`
	IsSend           bool    `json:"is_send"`
	ErrorMessage     *string `json:"error_message"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}
