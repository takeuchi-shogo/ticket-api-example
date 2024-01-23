package models

type RegisterEmails struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Token     string `json:"token"`
	PinCode   string `json:"pin_code"`
	IsValid   bool   `json:"is_valid"`
	IsSend    bool   `json:"is_send"`
	ExpireAt  int64  `json:"expire_at"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type RegisterEmailsResponse struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	// Token     string `json:"token"`
	// PinCode   string `json:"pin_code"`
	// IsValid   bool   `json:"is_valid"`
	// IsSend    bool   `json:"is_send"`
	// ExpireAt  int64  `json:"expire_at"`
}

func (r *RegisterEmails) BuildForGet() *RegisterEmailsResponse {
	return &RegisterEmailsResponse{
		ID:    r.ID,
		Email: r.Email,
	}
}
