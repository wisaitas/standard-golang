package main

import (
	_ "github.com/wisaitas/standard-golang/internal/standard-service/docs"
	"github.com/wisaitas/standard-golang/internal/standard-service/initial"
)

// @title Standard Service
// @version 1.0
// @description Standard Service
// @host localhost:8082
// @BasePath /api/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	initial.InitializeApp()
}
