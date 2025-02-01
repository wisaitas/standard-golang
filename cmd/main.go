package main

import (
	"github.com/wisaitas/standard-golang/internal/configs"
)

func init() {
	configs.LoadEnv()
}

func main() {
	app := configs.InitializeApp()

	app.SetupRoutes()

	app.Run()
}
