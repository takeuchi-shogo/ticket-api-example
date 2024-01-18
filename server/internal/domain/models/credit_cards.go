package models

type CreditCards struct {
	ID         int    `json:"id" bun:",pk,autoincrement"`
	UserID     int    `json:"user_id"`
	CustomerID string `json:"customer_id"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}

type CreditCardsReponse struct {
	Brand    string `json:"brand"`
	ExpMonth string `json:"exp_month"`
	ExpYear  string `json:"exp_year"`
	Last4    string `json:"last4"`
}
