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

func (u *UserRepository) FindByID(db bun.IDB, id int) (*models.Users, error) {
	user := &models.Users{}
	err := db.NewSelect().Model(user).Where("id = ?", id).Scan(context.Background())
	if err != nil {
		return &models.Users{}, errors.New("user is not found")
	}
	return user, nil
}

func (u *UserRepository) FindByEmail(db bun.IDB, email string) (*models.Users, error) {
	user := &models.Users{}
	fmt.Println(email)
	ctx := context.Background()
	err := db.NewSelect().Model(user).Where("email = ?", email).Scan(ctx)
	if err != nil {
		return &models.Users{}, errors.New("user is not found")
	}
	return user, nil
}

func (u *UserRepository) Create(db bun.IDB, user *models.Users) (*models.Users, error) {
	ctx := context.Background()

	user.CreatedAt = time.Now().Unix()
	user.UpdatedAt = time.Now().Unix()
	user.DeletedAt = nil

	if err := user.Validate(); err != nil {
		return &models.Users{}, err
	}

	_, err := db.NewInsert().Model(user).Exec(ctx)
	if err != nil {
		return &models.Users{}, err
	}
	return user, nil
}

func (u *UserRepository) Save(db bun.IDB, user *models.Users) (*models.Users, error) {

	user.UpdatedAt = time.Now().Unix()

	_, err := db.NewUpdate().
		Model(user).
		WherePK().
		Exec(context.Background())

	return user, err
}
