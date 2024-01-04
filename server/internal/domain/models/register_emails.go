package models

type RegisterEmails struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Token     string `json:"token"`
	PinCode   string `json:"pin_code"`
	IsValid   bool   `json:"is_valid"`
	ExpireAt  int64  `json:"expire_at"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}
