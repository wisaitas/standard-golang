package initial

import (
	"github.com/wisaitas/standard-golang/internal/handlers"
)

type handler struct {
	userHandler        handlers.UserHandler
	authHandler        handlers.AuthHandler
	provinceHandler    handlers.ProvinceHandler
	districtHandler    handlers.DistrictHandler
	subDistrictHandler handlers.SubDistrictHandler
}

func newHandler(service *service) *handler {
	return &handler{
		userHandler:        handlers.NewUserHandler(service.userService),
		authHandler:        handlers.NewAuthHandler(service.authService),
		provinceHandler:    handlers.NewProvinceHandler(service.provinceService),
		districtHandler:    handlers.NewDistrictHandler(service.districtService),
		subDistrictHandler: handlers.NewSubDistrictHandler(service.subDistrictService),
	}
}
