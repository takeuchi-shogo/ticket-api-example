package infrastructure

import (
	"log"
	"time"

	cron "github.com/go-co-op/gocron/v2"
	"github.com/takeuchi-shogo/ticket-api/config"
	"github.com/takeuchi-shogo/ticket-api/internal/adapters/controllers/tasks"
)

type Cron struct {
	Config      config.ServerConfig
	Controllers CronControllers
	Cron        cron.Scheduler
}

type CronControllers struct {
	draw tasks.DrawsController
}

func NewCron(
	c config.ServerConfig,
	draw tasks.DrawsController,
) *Cron {

	cron, err := cron.NewScheduler()
	if err != nil {
		panic(err)
	}

	controllers := CronControllers{
		draw: draw,
	}

	s := &Cron{
		Config:      c,
		Controllers: controllers,
		Cron:        cron,
	}

	s.newJob()

	return s
}

func (c *Cron) newJob() {

	_, err := c.Cron.NewJob(
		cron.DurationJob(
			time.Second*5,
		),
		cron.NewTask(
			func() {
				log.Println("test")
				c.Controllers.draw.Start()
			},
		),
	)
	_, _ = c.Cron.NewJob(
		cron.DurationJob(
			time.Second*6,
		),
		cron.NewTask(
			func() {
				log.Println("test2")
			},
		),
	)
	if err != nil {
		log.Fatal(err)
	}
}

func (c *Cron) setTasks() {}

func (c *Cron) Run() {
	// 本番環境のみ作動させる
	if c.Config.AppEnvironment != "production" {
		return
	}
	c.Cron.Start()
}
