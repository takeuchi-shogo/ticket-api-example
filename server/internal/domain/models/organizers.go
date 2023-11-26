package models

type Organizers struct {
	ID           int     `json:"id"`
	DisplayName  string  `json:"display_name"`
	ScreenName   string  `json:"screen_name"`
	Email        string  `json:"email"`
	Password     string  `json:"password"`
	CompanyUrl   *string `json:"company_url"`
	ContactTel   string  `json:"contact_tel"`
	ContactEmail string  `json:"contact_email"`

	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	DeletedAt *int64 `json:"deleted_at"`
}

type OrganizersResponse struct {
	ID           int    `json:"id"`
	DisplayName  string `json:"display_name"`
	ScreenName   string `json:"screen_name"`
	CompanyUrl   string `json:"company_url"`
	ContactTel   string `json:"contact_tel"`
	ContactEmail string `json:"contact_email"`
}
