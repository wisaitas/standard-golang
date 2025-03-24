package initial

import (
	authService "github.com/wisaitas/standard-golang/internal/services/auth"
	districtService "github.com/wisaitas/standard-golang/internal/services/district"
	provinceService "github.com/wisaitas/standard-golang/internal/services/province"
	subDistrictService "github.com/wisaitas/standard-golang/internal/services/sub-district"
	userService "github.com/wisaitas/standard-golang/internal/services/user"
)

type Service struct {
	UserService        userService.UserService
	AuthService        authService.AuthService
	ProvinceService    provinceService.ProvinceService
	DistrictService    districtService.DistrictService
	SubDistrictService subDistrictService.SubDistrictService
}

func NewService(repos *Repository, utils *Util) *Service {
	return &Service{
		UserService: userService.NewUserService(
			userService.NewRead(repos.UserRepository, utils.RedisUtil),
			userService.NewCreate(repos.UserRepository, utils.RedisUtil),
			userService.NewUpdate(repos.UserRepository, repos.UserHistoryRepository, utils.TransactionUtil, utils.RedisUtil),
			userService.NewDelete(repos.UserRepository, utils.RedisUtil),
			userService.NewTransaction(repos.UserRepository, utils.RedisUtil),
		),
		AuthService: authService.NewAuthService(repos.UserRepository, repos.UserHistoryRepository, utils.TransactionUtil, utils.RedisUtil, utils.BcryptUtil),
		ProvinceService: provinceService.NewProvinceService(
			provinceService.NewRead(repos.ProvinceRepository, utils.RedisUtil),
		),
		DistrictService: districtService.NewDistrictService(
			districtService.NewRead(repos.DistrictRepository, utils.RedisUtil),
		),
		SubDistrictService: subDistrictService.NewSubDistrictService(
			subDistrictService.NewRead(repos.SubDistrictRepository, utils.RedisUtil),
		),
	}
}
