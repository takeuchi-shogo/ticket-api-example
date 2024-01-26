package models

type Reviews struct {
	ID          int     `json:"id"`
	UserID      int     `json:"user_id"`
	VenueID     *int    `json:"venue_id"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Stars       float64 `json:"stars"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
	DeletedAt int64 `json:"deleted_at"`
}

type ReviewsResponse struct {
	ID          int     `json:"id"`
	UserID      int     `json:"user_id"`
	VenueID     *int    `json:"venue_id"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Stars       float64 `json:"stars"`

	User  *UsersResponse  `json:"user"`
	Venue *VenuesResponse `json:"venue"`

	CreatedAt int64 `json:"created_at"`
}
