package database

import (
	"context"
	"errors"
	"time"

	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
	"github.com/uptrace/bun"
)

type eventRepository struct{}

func NewEventRepository() usecase.EventUsecase {
	return &eventRepository{}
}

func (e *eventRepository) CountEventByArtistID(db bun.IDB, artistID int) (int, error) {
	count, err := db.NewSelect().Model((*models.Events)(nil)).Count(context.Background())
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (e *eventRepository) Find(db bun.IDB) ([]*models.Events, error) {

	events := []*models.Events{}

	ctx := context.Background()
	_ = db.NewSelect().Model(&events).Scan(ctx)
	if len(events) <= 0 {
		return []*models.Events{}, errors.New("events is not found")
	}
	return events, nil
}

func (e *eventRepository) FindByID(db bun.IDB, id int) (*models.Events, error) {
	event := &models.Events{}

	ctx := context.Background()

	_ = db.NewSelect().Model(event).Where("id = ?", id).Scan(ctx)
	if event.ID <= 0 {
		return &models.Events{}, errors.New("event is not found")
	}
	return event, nil
}

func (e *eventRepository) FindByArtistID(db bun.IDB, artistID int) ([]*models.Events, error) {
	events := []*models.Events{}

	ctx := context.Background()
	_ = db.NewSelect().Model(&events).Where("artist_id = ?", artistID).Scan(ctx)
	if len(events) <= 0 {
		return []*models.Events{}, errors.New("events is not found")
	}
	return events, nil
}

func (e *eventRepository) Create(db bun.IDB, event *models.Events) (*models.Events, error) {
	ctx := context.Background()

	event.CreatedAt = time.Now().Unix()
	event.UpdatedAt = time.Now().Unix()

	_, err := db.NewInsert().Model(event).Exec(ctx)
	if err != nil {
		return &models.Events{}, err
	}
	return event, nil
}
