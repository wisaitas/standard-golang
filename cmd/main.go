package main

import (
	"github.com/wisaitas/standard-golang/internal/initial"
	"github.com/wisaitas/standard-golang/internal/scripts"
)

func main() {
	scripts.HandleArgument()

	initial.InitializeApp()
}
