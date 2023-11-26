package models

type Events struct {
	ID                int     `json:"id"`
	OrganizerID       int     `json:"organizer_id"`
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
	VenueID           *int    `json:"venue_id"`
	Title             string  `json:"title"`
	PerformancePeriod string  `json:"performance_period"`
	EventType         string  `json:"event_type"`
	ShowTime          int64   `json:"show_time"`
	OpeningTime       int64   `json:"opening_time"`
	Description       *string `json:"description"`
	Note              *string `json:"note"`
	IsPrivate         bool    `json:"is_private"`
}