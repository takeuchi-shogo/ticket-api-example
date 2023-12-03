package database

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
	"github.com/uptrace/bun"
)

type UserRepository struct{}

func NewUserRepository() usecase.UserUsecase {
	return &UserRepository{}
}

func (u *UserRepository) FindByID(id int) (*models.Users, error) {
	return &models.Users{
		ID: uint64(id),
	}, nil
}

func (u *UserRepository) FindByEmail(db *bun.DB, email string) (*models.Users, error) {
	user := &models.Users{}
	fmt.Println(email)
	ctx := context.Background()
	err := db.NewSelect().Model(user).Where("email = ?", email).Scan(ctx)
	if err != nil {
		return &models.Users{}, errors.New("user is not found")
	}
	return user, nil
}

func (u *UserRepository) Create(db *bun.DB, user *models.Users) (*models.Users, error) {
	ctx := context.Background()

	user.CreatedAt = time.Now().Unix()
	user.UpdatedAt = time.Now().Unix()

	_, err := db.NewInsert().Model(user).Exec(ctx)
	if err != nil {
		return &models.Users{}, err
	}
	return user, nil
}
