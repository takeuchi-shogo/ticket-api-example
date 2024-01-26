package database

import "github.com/takeuchi-shogo/ticket-api/internal/usecase"

type eventHasArtistRepository struct{}

func NewEventHasArtistRepository() usecase.EventHasArtistUsecase {
	return &eventHasArtistRepository{}
}
