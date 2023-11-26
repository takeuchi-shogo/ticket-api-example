package models

type Administrators struct {
	ID          int    `json:"id"`
	DisplayName string `json:"display_name"`
	ScreenName  string `json:"screen_name"`
	Password    string `json:"password"`
	Role        string `json:"role"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
	DeletedAt int64 `json:"deleted_at"`
}

type AdministratorsResponse struct {
	ID          int    `json:"id"`
	DisplayName string `json:"display_name"`
	ScreenName  string `json:"screen_name"`
	Role        string `json:"role"`
}
