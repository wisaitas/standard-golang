package initial

import (
	authService "github.com/wisaitas/standard-golang/internal/standard-service/service/auth"
	districtService "github.com/wisaitas/standard-golang/internal/standard-service/service/district"
	provinceService "github.com/wisaitas/standard-golang/internal/standard-service/service/province"
	subDistrictService "github.com/wisaitas/standard-golang/internal/standard-service/service/sub-district"
	userService "github.com/wisaitas/standard-golang/internal/standard-service/service/user"
)

type service struct {
	userService        userService.UserService
	authService        authService.AuthService
	provinceService    provinceService.ProvinceService
	districtService    districtService.DistrictService
	subDistrictService subDistrictService.SubDistrictService
}

func newService(repo *repository, sharePkg *sharePkg) *service {
	return &service{
		userService: userService.NewUserService(
			userService.NewGet(repo.userRepository, sharePkg.redis),
			userService.NewPost(repo.userRepository, sharePkg.redis),
			userService.NewUpdate(repo.userRepository, repo.userHistoryRepository, sharePkg.redis, sharePkg.transactionManager),
			userService.NewDelete(repo.userRepository, sharePkg.redis),
		),
		authService: authService.NewAuthService(
			repo.userRepository,
			repo.userHistoryRepository,
			sharePkg.redis,
			sharePkg.bcrypt,
			sharePkg.jwt,
			sharePkg.transactionManager,
		),
		provinceService: provinceService.NewProvinceService(
			provinceService.NewGet(repo.provinceRepository, sharePkg.redis),
		),
		districtService: districtService.NewDistrictService(
			districtService.NewGet(repo.districtRepository, sharePkg.redis),
		),
		subDistrictService: subDistrictService.NewSubDistrictService(
			subDistrictService.NewGet(repo.subDistrictRepository, sharePkg.redis),
		),
	}
}
