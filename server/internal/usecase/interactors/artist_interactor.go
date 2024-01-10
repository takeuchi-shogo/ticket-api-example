package interactors

import (
	"net/http"

	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase/services"
	"github.com/uptrace/bun"
)

type artistInteractor struct {
	db     usecase.DBUsecase
	artist usecase.ArtistUsecase
	event  usecase.EventUsecase
}

func NewArtistInteractor(
	db usecase.DBUsecase,
	artist usecase.ArtistUsecase,
	event usecase.EventUsecase,
) services.ArtistService {
	return &artistInteractor{
		db:     db,
		artist: artist,
		event:  event,
	}
}

func (a *artistInteractor) Get(id int) (*models.ArtistsResponse, *usecase.ResultStatus) {

	db, _ := a.db.Connect()

	artist, err := a.artist.FindByID(db, id)
	if err != nil {
		return &models.ArtistsResponse{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	builtArtist, err := a.build(db, artist)
	if err != nil {
		return &models.ArtistsResponse{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	return builtArtist, usecase.NewResultStatus(http.StatusOK, nil)
}

func (a *artistInteractor) GetList() ([]*models.ArtistsResponse, *usecase.ResultStatus) {

	db, _ := a.db.Connect()

	artists, err := a.artist.Find(db)
	if err != nil {
		return []*models.ArtistsResponse{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	builtArtists := make([]*models.ArtistsResponse, len(artists))

	for i, artist := range artists {
		builtArtists[i] = artist.BuildForGet()
	}

	return builtArtists, usecase.NewResultStatus(http.StatusOK, nil)
}

func (a *artistInteractor) build(db bun.IDB, artist *models.Artists) (*models.ArtistsResponse, error) {

	eventCnt, err := a.event.CountEventByArtistID(db, artist.ID)
	if err != nil {
		return &models.ArtistsResponse{}, err
	}
	builtArtist := artist.BuildForGet()
	builtArtist.EventCnt = eventCnt

	return builtArtist, nil
}

func (a *artistInteractor) Create(artist *models.Artists) (*models.ArtistsResponse, *usecase.ResultStatus) {

	db, _ := a.db.Connect()

	newArtist, err := a.artist.Create(db, artist)
	if err != nil {
		return &models.ArtistsResponse{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	return newArtist.BuildForGet(), usecase.NewResultStatus(http.StatusOK, nil)
}
