package main

import (
	"log"

	"github.com/takeuchi-shogo/ticket-api/config"
	"github.com/takeuchi-shogo/ticket-api/internal/adapters/gateways"
	"github.com/takeuchi-shogo/ticket-api/internal/adapters/gateways/database"
	"github.com/takeuchi-shogo/ticket-api/internal/infrastructure"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase/interactors/tasks"
)

// 抽選処理を実行する
func main() {
	config := config.NewConfig()
	db := infrastructure.NewDB(config)

	dbRepo := gateways.NewDBGateway(db)
	ticketRepo := database.NewTicketRepository()
	userBookTicketRepo := database.NewUserBookTicketRepository()
	userHasTicketRepo := database.NewUserHasTicketRepository()

	interactor := tasks.NewDrawInteractor(
		dbRepo,
		ticketRepo,
		userBookTicketRepo,
		userHasTicketRepo,
	)

	err := interactor.StartDrawing()
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("success")
}
