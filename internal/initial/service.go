package initial

import (
	authService "github.com/wisaitas/standard-golang/internal/services/auth"
	districtService "github.com/wisaitas/standard-golang/internal/services/district"
	provinceService "github.com/wisaitas/standard-golang/internal/services/province"
	subDistrictService "github.com/wisaitas/standard-golang/internal/services/sub-district"
	userService "github.com/wisaitas/standard-golang/internal/services/user"
)

type service struct {
	userService        userService.UserService
	authService        authService.AuthService
	provinceService    provinceService.ProvinceService
	districtService    districtService.DistrictService
	subDistrictService subDistrictService.SubDistrictService
}

func newService(repo *repository, util *util) *service {
	return &service{
		userService: userService.NewUserService(
			userService.NewRead(repo.userRepository, util.redisUtil),
			userService.NewCreate(repo.userRepository, util.redisUtil),
			userService.NewUpdate(repo.userRepository, repo.userHistoryRepository, util.transactionUtil, util.redisUtil),
			userService.NewDelete(repo.userRepository, util.redisUtil),
			userService.NewTransaction(repo.userRepository, util.redisUtil),
		),
		authService: authService.NewAuthService(repo.userRepository, repo.userHistoryRepository, util.transactionUtil, util.redisUtil, util.bcryptUtil),
		provinceService: provinceService.NewProvinceService(
			provinceService.NewRead(repo.provinceRepository, util.redisUtil),
		),
		districtService: districtService.NewDistrictService(
			districtService.NewRead(repo.districtRepository, util.redisUtil),
		),
		subDistrictService: subDistrictService.NewSubDistrictService(
			subDistrictService.NewGet(repo.subDistrictRepository, util.redisUtil),
		),
	}
}
