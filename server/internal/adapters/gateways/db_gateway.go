package gateways

import (
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
	"github.com/uptrace/bun"
)

type DB interface {
	Connect() *bun.DB
	Transaction() (bun.Tx, error)
}

type DBGateway struct {
	DB DB
}

func NewDBGateway(db DB) usecase.DBUsecase {
	return &DBGateway{DB: db}
}

func (g *DBGateway) Connect() *bun.DB {
	return g.DB.Connect()
}

func (g *DBGateway) Transaction() (bun.Tx, error) {
	return g.DB.Transaction()
}
