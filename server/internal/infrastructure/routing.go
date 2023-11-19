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
	db   *DB
	Gin  *gin.Engine
	Port string
}

type Controllers struct {
	auth  controllers.AuthController
	event controllers.EventController
	user  controllers.UserController
}

func NewControllers(
	a controllers.AuthController,
	e controllers.EventController,
	u controllers.UserController,
) Controllers {
	return Controllers{
		auth:  a,
		event: e,
		user:  u,
	}
}

func NewRouting(config config.ServerConfig, db *DB, c Controllers) *Routing {

	if config.AppEnvironment == "development" {
		gin.SetMode(gin.DebugMode)
		fmt.Println("=========== development ===============")
	}

	r := &Routing{
		db:   db,
		Gin:  gin.Default(),
		Port: fmt.Sprintf(":%d", config.Port),
	}

	r.Gin.POST("/signin", func(ctx *gin.Context) { c.auth.Signin(ctx) })
	r.Gin.POST("/signup", func(ctx *gin.Context) { c.auth.Signup(ctx) })

	r.Gin.GET("/events/:id", func(ctx *gin.Context) { c.event.Get(ctx) })

	r.Gin.GET("/users/:id", func(ctx *gin.Context) { c.user.Get(ctx) })

	return r
}

func (r *Routing) Run() {
	// fmt.Printf("finasldn: %+v", r)
	s := &http.Server{
		Addr:           r.Port,
		Handler:        r.Gin,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
