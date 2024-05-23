package tasks

import (
	"log"

	"github.com/takeuchi-shogo/ticket-api/internal/usecase/services/tasks"
)

type DrawsController interface {
	Start()
}

type drawsController struct {
	draw tasks.DrawService
}

func NewDrawsController() DrawsController {
	return &drawsController{}
}

func (d *drawsController) Start() {
	if err := d.draw.StartDrawing(); err != nil {
		log.Println(err)
		return
	}
	log.Println("success")
}
