package bootstrap

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/takeuchi-shogo/ticket-api/internal/infrastructure"
	"go.uber.org/fx"
)

var rootCmd = &cobra.Command{
	Use:   "clean-gin",
	Short: "Clean architecture using gin framework",
	Long: `
█▀▀ █░░ █▀▀ ▄▀█ █▄░█ ▄▄ █▀▀ █ █▄░█
█▄▄ █▄▄ ██▄ █▀█ █░▀█ ░░ █▄█ █ █░▀█      
                                         		
This is a command runner or cli for api architecture in golang. 
Using this we can use underlying dependency injection container for running scripts. 
Main advantage is that, we can use same services, repositories, infrastructure present in the application itself`,
	TraverseChildren: true,
}

type App struct {
	// *cobra.Command
	FxApp *fx.App
}

func NewApp() App {
	cmd := App{
		// Command: rootCmd,
	}

	cmd.FxApp = fx.New(
		CommonModule,
		fx.Invoke(
			NewServer,
		),
	)

	return cmd
}

func NewServer(lc fx.Lifecycle, r *infrastructure.Routing, c *infrastructure.Cron) {

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// ln, err := net.Listen("tcp", srv.Addr) // the web server starts listening on 8080
			// if err != nil {
			// 	fmt.Println("[My Demo] Failed to start HTTP Server at")
			// 	return err
			// }
			// go srv.Serve(ln) // process an incoming request in a go routine
			go r.Run()
			go c.Run()
			fmt.Println("[My Demo]Succeeded to start HTTP Server at")
			return nil

		},
		OnStop: func(ctx context.Context) error {
			// srv.Shutdown(ctx) // stop the web server
			fmt.Println("[My Demo] HTTP Server is stopped")
			return nil
		},
	})
}

// var RootApp = NewApp()
