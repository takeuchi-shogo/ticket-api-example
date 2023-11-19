package main

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func addGroup(r *gin.Engine) {
	api := r.Group("/api")
	{
		admin := api.Group(("/admin"))
		{
			admin.GET("/", adminFunction)
		}
	}
}

func adminFunction(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"admin-function": "admin function content."})
}

func NewServer(lc fx.Lifecycle) *gin.Engine {
	router := gin.Default()
	addGroup(router)

	srv := &http.Server{Addr: ":8080", Handler: router}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr) // the web server starts listening on 8080
			if err != nil {
				fmt.Println("[My Demo] Failed to start HTTP Server at", srv.Addr)
				return err
			}
			go srv.Serve(ln) // process an incoming request in a go routine
			fmt.Println("[My Demo]Succeeded to start HTTP Server at", srv.Addr)
			return nil

		},
		OnStop: func(ctx context.Context) error {
			srv.Shutdown(ctx) // stop the web server
			fmt.Println("[My Demo] HTTP Server is stopped")
			return nil
		},
	})

	return router
}

func main() {
	app := fx.New(
		fx.Provide(
			NewServer,
		),
		fx.Invoke(func(*gin.Engine) {}),
	)
	app.Run()
}
