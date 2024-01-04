package database

import (
	"context"
	"time"

	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
	"github.com/uptrace/bun"
)

type RegisterEmailRepository struct {
}

func NewRegisterEmailRepository() usecase.RegisterEmailUsecase {
	return &RegisterEmailRepository{}
}

func (t *RegisterEmailRepository) FindByEmail(db bun.IDB, email string) (*models.RegisterEmails, error) {
	return &models.RegisterEmails{}, nil
}

func (t *RegisterEmailRepository) FindByPinCode(db bun.IDB, pinCode string) (*models.RegisterEmails, error) {

	registerEmail := &models.RegisterEmails{}
	err := db.NewSelect().
		Model(registerEmail).
		Where("").
		Scan(context.Background())

	return registerEmail, err
}

func (t *RegisterEmailRepository) Create(db bun.IDB, registerEmail *models.RegisterEmails) (*models.RegisterEmails, error) {

	registerEmail.IsValid = false
	registerEmail.ExpireAt = time.Now().Add(1 * time.Hour).Unix()
	registerEmail.CreatedAt = time.Now().Unix()
	registerEmail.UpdatedAt = time.Now().Unix()

	_, err := db.NewInsert().
		Model(registerEmail).
		Exec(context.Background())

	return registerEmail, err
}
