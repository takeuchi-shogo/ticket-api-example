package database

import (
	"context"
	"time"

	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
	"github.com/uptrace/bun"
)

type ArtistRepository struct{}

func NewArtistRepository() usecase.ArtistUsecase { return &ArtistRepository{} }

func (a *ArtistRepository) Find(db bun.IDB) ([]*models.Artists, error) {
	artists := []*models.Artists{}
	if err := db.NewSelect().Model(&artists).Scan(context.Background()); err != nil {
		return []*models.Artists{}, err
	}
	return artists, nil
}

func (a *ArtistRepository) FindByID(db bun.IDB, id int) (*models.Artists, error) {
	artist := &models.Artists{}
	if err := db.NewSelect().
		Model(artist).
		Where("id = ?", id).
		Scan(context.Background()); err != nil {
		return &models.Artists{}, err
	}

	return artist, nil
}

func (a *ArtistRepository) Create(db bun.IDB, artist *models.Artists) (*models.Artists, error) {
	ctx := context.Background()

	artist.CreatedAt = time.Now().Unix()
	artist.UpdatedAt = time.Now().Unix()

	_, err := db.NewInsert().Model(artist).Exec(ctx)
	if err != nil {
		return &models.Artists{}, err
	}
	return artist, nil
}
