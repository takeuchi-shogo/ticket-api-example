package models

type Venues struct {
	ID                 int     `json:"id"`
	Name               string  `json:"name"`
	PostCode           string  `json:"post_code"`
	Address            string  `json:"address"`
	HowToAccess        *string `json:"how_to_access"`
	Capacity           int     `json:"capacity"`
	ParkingSpace       int     `json:"parking_space"`
	ParkingDescription *string `json:"parking_description"`
	SiteUrl            *string `json:"site_url"`
	ContactTel         *string `json:"contact_tel"`
	ContactEmail       *string `json:"contact_email"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}

type VenuesResponse struct {
	ID                 int     `json:"id"`
	Name               string  `json:"name"`
	PostCode           string  `json:"post_code"`
	Address            string  `json:"address"`
	HowToAccess        *string `json:"how_to_access"`
	Capacity           int     `json:"capacity"`
	ParkingSpace       int     `json:"parking_space"`
	ParkingDescription *string `json:"parking_description"`
	SiteUrl            *string `json:"site_url"`
	ContactTel         *string `json:"contact_tel"`
	ContactEmail       *string `json:"contact_email"`
}
