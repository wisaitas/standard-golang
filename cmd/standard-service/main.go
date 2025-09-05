package main

import (
	"github.com/wisaitas/standard-golang/internal/standard-service/initial"
)

func main() {
	app := initial.InitializeApp()

	app.Run()

	app.Cleanup()
}
