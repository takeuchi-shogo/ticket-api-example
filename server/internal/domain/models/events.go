package models

type Events struct {
	ID                int     `json:"id" bun:",pk,autoincrement"`
	OrganizerID       int     `json:"organizer_id"`
	ArtistID          int     `json:"artist_id"`
	VenueID           *int    `json:"venue_id"`
	Title             string  `json:"title"`
	PerformancePeriod string  `json:"performance_period"`
	EventType         string  `json:"event_type"`
	ShowTime          int64   `json:"show_time"`
	OpeningTime       int64   `json:"opening_time"`
	Description       *string `json:"description"`
	Note              *string `json:"note"`
	IsPrivate         bool    `json:"is_private"`

	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	DeletedAt *int64 `json:"deleted_at"`
}

type EventsReponse struct {
	ID                int     `json:"id"`
	OrganizerID       int     `json:"organizer_id"`
	ArtistID          int     `json:"artist_id"`
	VenueID           *int    `json:"venue_id"`
	Title             string  `json:"title"`
	PerformancePeriod string  `json:"performance_period"`
	EventType         string  `json:"event_type"`
	EventCategory     string  `json:"event_category"`
	ShowTime          int64   `json:"show_time"`
	OpeningTime       int64   `json:"opening_time"`
	Description       *string `json:"description"`
	Note              *string `json:"note"`
	IsPrivate         bool    `json:"is_private"`

	Artists []*ArtistsResponse `json:"artists,omitempty"`
}

type EventInteractorResponse struct {
	Total  int              `json:"total"`
	Events []*EventsReponse `json:"events"`
}

type EventType struct {
}

func (e *Events) BuildFor() *EventsReponse {
	return &EventsReponse{
		ID:                e.ID,
		OrganizerID:       e.OrganizerID,
		ArtistID:          e.ArtistID,
		VenueID:           e.VenueID,
		Title:             e.Title,
		PerformancePeriod: e.PerformancePeriod,
		EventType:         e.EventType,
		ShowTime:          e.ShowTime,
		OpeningTime:       e.OpeningTime,
		Description:       e.Description,
		Note:              e.Note,
		IsPrivate:         e.IsPrivate,
	}
}
