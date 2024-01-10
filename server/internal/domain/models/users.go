package models

type Users struct {
	ID          uint64 `json:"id" bun:",pk"`
	DisplayName string `json:"display_name"`
	ScreenName  string `json:"screen_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
	DeletedAt int64 `json:"deleted_at"`
}

type UsersResponse struct {
	ID          uint64 `json:"id"`
	DisplayName string `json:"display_name"`
	ScreenName  string `json:"screen_name"`
	Email       string `json:"email"`
}

type MeInteractorResponse struct {
	User  *UsersResponse `json:"user"`
	Token string         `json:"token"`
}

func (u *Users) BuildForGet() *UsersResponse {
	return &UsersResponse{
		ID:          u.ID,
		DisplayName: u.DisplayName,
		ScreenName:  u.ScreenName,
		Email:       u.Email,
	}
}
