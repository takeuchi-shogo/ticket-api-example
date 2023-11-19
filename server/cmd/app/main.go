package main

import "github.com/takeuchi-shogo/ticket-api/bootstrap"

func main() {
	// bootstrap.RootApp.Execute()
	app := bootstrap.NewApp()
	app.FxApp.Run()
}
