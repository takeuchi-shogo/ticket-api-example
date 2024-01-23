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

func (t *RegisterEmailRepository) FindByNotSend(db bun.IDB) ([]*models.RegisterEmails, error) {
	emails := []*models.RegisterEmails{}
	err := db.NewSelect().
		Model(emails).
		Where("is_send = ?", false).
		Scan(context.Background())

	return emails, err
}

func (t *RegisterEmailRepository) FindByEmail(db bun.IDB, email string) (*models.RegisterEmails, error) {
	return &models.RegisterEmails{}, nil
}

func (t *RegisterEmailRepository) FindByToken(db bun.IDB, token string) (*models.RegisterEmails, error) {
	registerEmail := &models.RegisterEmails{}
	err := db.NewSelect().
		Model(registerEmail).
		Where("token = ?", token).
		Scan(context.Background())

	return registerEmail, err
}

func (t *RegisterEmailRepository) FindByPinCode(db bun.IDB, pinCode string) (*models.RegisterEmails, error) {
	registerEmail := &models.RegisterEmails{}
	err := db.NewSelect().
		Model(registerEmail).
		Where("pin_code = ?", pinCode).
		Scan(context.Background())

	return registerEmail, err
}

func (t *RegisterEmailRepository) Create(db bun.IDB, registerEmail *models.RegisterEmails) (*models.RegisterEmails, error) {

	registerEmail.IsValid = true
	registerEmail.IsSend = false
	registerEmail.ExpireAt = time.Now().Add(1 * time.Hour).Unix() // 1時間の有効期限
	registerEmail.CreatedAt = time.Now().Unix()
	registerEmail.UpdatedAt = time.Now().Unix()

	_, err := db.NewInsert().
		Model(registerEmail).
		Exec(context.Background())

	return registerEmail, err
}
