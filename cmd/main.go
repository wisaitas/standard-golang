package main

import "github.com/wisaitas/standard-golang/internal/initial"

func main() {
	app := initial.InitializeApp()

	app.SetupMiddlewares()

	app.SetupRoutes()

	app.Run()
}
