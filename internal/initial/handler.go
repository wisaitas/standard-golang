package initial

import (
	"github.com/wisaitas/standard-golang/internal/handlers"
)

type Handlers struct {
	UserHandler        handlers.UserHandler
	AuthHandler        handlers.AuthHandler
	ProvinceHandler    handlers.ProvinceHandler
	DistrictHandler    handlers.DistrictHandler
	SubDistrictHandler handlers.SubDistrictHandler
}

func initializeHandlers(services *Services) *Handlers {
	return &Handlers{
		UserHandler:        *handlers.NewUserHandler(services.UserService),
		AuthHandler:        *handlers.NewAuthHandler(services.AuthService),
		ProvinceHandler:    *handlers.NewProvinceHandler(services.ProvinceService),
		DistrictHandler:    *handlers.NewDistrictHandler(services.DistrictService),
		SubDistrictHandler: *handlers.NewSubDistrictHandler(services.SubDistrictService),
	}
}
