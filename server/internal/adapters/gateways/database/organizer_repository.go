package database

import (
	"context"
	"errors"
	"time"

	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
	"github.com/uptrace/bun"
)

type organizerRepository struct{}

func NewOrganizerRepository() usecase.OrganizerUsecase {
	return &organizerRepository{}
}

func (o *organizerRepository) FindByID(db bun.IDB, id int) (*models.Organizers, error) {
	ctx := context.Background()

	organizer := &models.Organizers{}

	_ = db.NewSelect().Model(organizer).Where("id = ?", id).Scan(ctx)
	if organizer.ID <= 0 {
		return &models.Organizers{}, errors.New("organizer is not found")
	}

	return organizer, nil
}

func (o *organizerRepository) Create(db bun.IDB, organizer *models.Organizers) (*models.Organizers, error) {
	ctx := context.Background()

	organizer.CreatedAt = time.Now().Unix()
	organizer.UpdatedAt = time.Now().Unix()

	_, err := db.NewInsert().Model(organizer).Exec(ctx)
	if err != nil {
		return &models.Organizers{}, err
	}

	return organizer, nil
}
