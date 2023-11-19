package database

import "github.com/takeuchi-shogo/ticket-api/internal/domain/models"

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (u *UserRepository) FindByID(id int) (*models.Users, error) {
	return &models.Users{
		ID: id,
	}, nil
}
