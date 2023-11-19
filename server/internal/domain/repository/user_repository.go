package repository

import "github.com/takeuchi-shogo/ticket-api/internal/domain/models"

type IUserRepository interface {
	FindByID(id int) (*models.Users, error)
}
