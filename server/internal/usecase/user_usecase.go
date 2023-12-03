package usecase

import (
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/uptrace/bun"
)

type UserUsecase interface {
	FindByID(id int) (*models.Users, error)
	FindByEmail(db *bun.DB, email string) (*models.Users, error)
	Create(db *bun.DB, user *models.Users) (*models.Users, error)
}
