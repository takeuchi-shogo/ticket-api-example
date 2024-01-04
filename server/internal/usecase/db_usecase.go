package usecase

import "github.com/uptrace/bun"

type DBUsecase interface {
	Connect() (bun.Conn, error)
	Transaction() (bun.Tx, error)
}
