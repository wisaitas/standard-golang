package initial

import (
	authService "github.com/wisaitas/standard-golang/internal/standard-service/services/auth"
	districtService "github.com/wisaitas/standard-golang/internal/standard-service/services/district"
	provinceService "github.com/wisaitas/standard-golang/internal/standard-service/services/province"
	subDistrictService "github.com/wisaitas/standard-golang/internal/standard-service/services/sub-district"
	userService "github.com/wisaitas/standard-golang/internal/standard-service/services/user"
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
			userService.NewGet(repo.userRepository, util.redisUtil),
			userService.NewPost(repo.userRepository, util.redisUtil),
			userService.NewUpdate(repo.userRepository, repo.userHistoryRepository, util.transactionUtil, util.redisUtil),
			userService.NewDelete(repo.userRepository, util.redisUtil),
			userService.NewTransaction(repo.userRepository, util.redisUtil),
		),
		authService: authService.NewAuthService(repo.userRepository, repo.userHistoryRepository, util.transactionUtil, util.redisUtil, util.bcryptUtil),
		provinceService: provinceService.NewProvinceService(
			provinceService.NewGet(repo.provinceRepository, util.redisUtil),
		),
		districtService: districtService.NewDistrictService(
			districtService.NewGet(repo.districtRepository, util.redisUtil),
		),
		subDistrictService: subDistrictService.NewSubDistrictService(
			subDistrictService.NewGet(repo.subDistrictRepository, util.redisUtil),
		),
	}
}
