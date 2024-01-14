package models

type PaymentByCreditCards struct {
	ID               int
	UserBookTicketID int
	PaymentID        string
	IsValid          bool
	ExpireAt         int64

	CreatedAt int64
	UpdatedAt int64
}
