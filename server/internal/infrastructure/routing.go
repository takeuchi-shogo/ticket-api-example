package infrastructure

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/takeuchi-shogo/ticket-api/config"
	"github.com/takeuchi-shogo/ticket-api/internal/adapters/controllers"
)

const (
	ApiVersion = "/v1"

	HealthCheckRoot = "/healthChecks"

	// API Group
	ProductGroup = "/product"

	// API Routing
	UserRoot = "/users"
)

type Routing struct {
	// db   *DB
	Gin  *gin.Engine
	Port string
}

type Controllers struct {
	admin     controllers.AdministratorsController
	auth      controllers.AuthController
	event     controllers.EventController
	organizer controllers.OrganizersController
	ticket    controllers.TicketsController
	user      controllers.UserController
}

func NewControllers(
	ad controllers.AdministratorsController,
	a controllers.AuthController,
	e controllers.EventController,
	o controllers.OrganizersController,
	t controllers.TicketsController,
	u controllers.UserController,
) Controllers {
	return Controllers{
		admin:     ad,
		auth:      a,
		event:     e,
		organizer: o,
		ticket:    t,
		user:      u,
	}
}

func NewRouting(config config.ServerConfig, c Controllers) *Routing {

	if config.AppEnvironment == "development" {
		gin.SetMode(gin.DebugMode)
		fmt.Println("=========== development ===============")
	}

	r := &Routing{
		Gin:  gin.Default(),
		Port: fmt.Sprintf(":%d", config.Port),
	}

	r.Gin.POST("/signin", func(ctx *gin.Context) { c.auth.Signin(ctx) })
	r.Gin.POST("/signup", func(ctx *gin.Context) { c.auth.Signup(ctx) })

	r.Gin.GET("/events/:id", func(ctx *gin.Context) { c.event.Get(ctx) })

	r.Gin.GET("/organizers/:id", func(ctx *gin.Context) { c.organizer.Get(ctx) })

	r.Gin.GET("/tickets/:id", func(ctx *gin.Context) { c.ticket.Get(ctx) })

	r.Gin.GET("/users/:id", func(ctx *gin.Context) { c.user.Get(ctx) })

	v1Admin := r.Gin.Group("/admin")
	{
		v1Admin.GET("/administrators/:id", func(ctx *gin.Context) { c.admin.Get(ctx) })
	}

	return r
}

func (r *Routing) Run() {
	s := &http.Server{
		Addr:           r.Port,
		Handler:        r.Gin,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()

	fmt.Printf("Running server: %s\n", r.Port)
}
