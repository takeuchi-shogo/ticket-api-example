package usecase

import (
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/uptrace/bun"
)

type UserUsecase interface {
	FindByID(db bun.IDB, userID int) (*models.Users, error)
	FindByEmail(db bun.IDB, email string) (*models.Users, error)
	Create(db bun.IDB, user *models.Users) (*models.Users, error)
	Save(db bun.IDB, user *models.Users) (*models.Users, error)
}
