package initial

import (
	handlerInternal "github.com/wisaitas/standard-golang/internal/standard-service/handler"
)

type handler struct {
	userHandler        handlerInternal.UserHandler
	authHandler        handlerInternal.AuthHandler
	provinceHandler    handlerInternal.ProvinceHandler
	districtHandler    handlerInternal.DistrictHandler
	subDistrictHandler handlerInternal.SubDistrictHandler
}

func newHandler(service *service) *handler {
	return &handler{
		userHandler:        handlerInternal.NewUserHandler(service.userService),
		authHandler:        handlerInternal.NewAuthHandler(service.authService),
		provinceHandler:    handlerInternal.NewProvinceHandler(service.provinceService),
		districtHandler:    handlerInternal.NewDistrictHandler(service.districtService),
		subDistrictHandler: handlerInternal.NewSubDistrictHandler(service.subDistrictService),
	}
}
