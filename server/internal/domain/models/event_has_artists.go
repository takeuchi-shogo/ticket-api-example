package models

type EventHasArtists struct {
	ID       int `json:"id"`
	EventID  int `json:"event_id"`
	ArtistID int `json:"artist_id"`

	CreatedAt int64 `json:"created_at"`
}
