package usecase

import "github.com/uptrace/bun"

type DBUsecase interface {
	Connect() *bun.DB
	Transaction() (bun.Tx, error)
}
