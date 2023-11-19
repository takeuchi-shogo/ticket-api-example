package services

import "github.com/takeuchi-shogo/ticket-api/internal/domain/models"

type UserService interface {
	Get(id int) (*models.Users, error)
}
