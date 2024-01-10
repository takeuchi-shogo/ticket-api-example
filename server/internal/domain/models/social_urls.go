package models

type SocialURLs struct {
	ID              int    `json:"id"`
	ArtistID        int    `json:"artist_id"`
	ApplicationName string `json:"application_name"`
	Url             string `json:"url"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}

type SocialURLsResponse struct {
	ID              int    `json:"id"`
	ArtistID        int    `json:"artist_id"`
	ApplicationName string `json:"application_name"`
	Url             string `json:"url"`
}
