package models

type Artists struct {
	ID          int     `json:"id"`
	DisplayName string  `json:"display_name"`
	ScreenName  string  `json:"screen_name"`
	Description *string `json:"description"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
	DeletedAt int64 `json:"deleted_at"`
}

type ArtistsResponse struct {
	ID          int    `json:"id"`
	DisplayName string `json:"display_name"`
	ScreenName  string `json:"screen_name"`
	Description string `json:"description"`

	EventCnt       int     `json:"event_cnt"`
	ReviewsCnt     int     `json:"reviews_cnt"`
	ReviewsRaiting float64 `json:"reviews_raiting"`
	// Reviews      []*ReviewsResponse `json:"reviews"`

	// SocialURLs []*SocialURLsResponse `json:"urls"`
}

func (a Artists) BuildForGet() *ArtistsResponse {
	return &ArtistsResponse{
		ID:          a.ID,
		DisplayName: a.DisplayName,
		ScreenName:  a.ScreenName,
		Description: *a.Description,
	}
}
