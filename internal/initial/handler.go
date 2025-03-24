package initial

import (
	"github.com/wisaitas/standard-golang/internal/handlers"
)

type Handler struct {
	UserHandler        handlers.UserHandler
	AuthHandler        handlers.AuthHandler
	ProvinceHandler    handlers.ProvinceHandler
	DistrictHandler    handlers.DistrictHandler
	SubDistrictHandler handlers.SubDistrictHandler
}

func NewHandler(services *Service) *Handler {
	return &Handler{
		UserHandler:        *handlers.NewUserHandler(services.UserService),
		AuthHandler:        *handlers.NewAuthHandler(services.AuthService),
		ProvinceHandler:    *handlers.NewProvinceHandler(services.ProvinceService),
		DistrictHandler:    *handlers.NewDistrictHandler(services.DistrictService),
		SubDistrictHandler: *handlers.NewSubDistrictHandler(services.SubDistrictService),
	}
}
