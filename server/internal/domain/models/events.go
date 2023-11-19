package models

type Events struct {
	ID                int    `json:"id"`
	Title             string `json:"title"`
	PerformancePeriod string `json:"performance_period"`
	Description       string `json:"description"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}
