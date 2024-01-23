package usecase

import (
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/uptrace/bun"
)

type RegisterEmailUsecase interface {
	FindByNotSend(db bun.IDB) ([]*models.RegisterEmails, error)
	FindByEmail(db bun.IDB, email string) (*models.RegisterEmails, error)
	FindByToken(db bun.IDB, token string) (*models.RegisterEmails, error)
	FindByPinCode(db bun.IDB, pinCode string) (*models.RegisterEmails, error)
	Create(db bun.IDB, registerEmail *models.RegisterEmails) (*models.RegisterEmails, error)
}
