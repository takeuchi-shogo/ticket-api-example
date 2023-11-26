package models

type Blogs struct {
	ID              int    `json:"id"`
	AdministratorID int    `json:"administrator_id"`
	Title           string `json:"title"`
	Content         string `json:"content"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
	DeletedAt int64 `json:"deleted_at"`
}

type BlogsResponse struct {
	ID              int    `json:"id"`
	AdministratorID int    `json:"administrator_id"`
	Title           string `json:"title"`
	Content         string `json:"content"`

	CreatedStr string `json:"created_str"`
	UpdatedStr string `json:"updated_str"`
}
