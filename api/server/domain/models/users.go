package models

type Users struct {
	ID          int    `json:"id"`
	DisplayName string `json:"display_name"`
	ScreenName  string `json:"screen_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
	DeletedAt int64 `json:"deleted_at"`
}
