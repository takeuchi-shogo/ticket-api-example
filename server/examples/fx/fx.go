package main

import (
	"fmt"

	"go.uber.org/fx"
)

type Usecase interface {
	Use()
}

type Repository interface {
	RepoPrint()
}

type usecase struct {
	repo Repository
}

func NewUsecase(r Repository) Usecase {
	return &usecase{
		repo: r,
	}
}

func (u usecase) Use() {
	u.repo.RepoPrint()
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r repository) RepoPrint() {
	fmt.Println("done!!!")
}

func main() {
	// Normal
	// repo := NewRepository()
	// usecase := NewUsecase(repo)
	// usecase.Use()

	app := fx.New(
		fx.Provide(
			NewUsecase,
			NewRepository,
		),
		fx.Invoke(func(u Usecase) {
			u.Use()
		}),
	)

	app.Done()
}
