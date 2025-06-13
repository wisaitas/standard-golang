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

func newService(repo *repository, lib *lib) *service {
	return &service{
		userService: userService.NewUserService(
			userService.NewGet(repo.userRepository, lib.redis),
			userService.NewPost(repo.userRepository, lib.redis),
			userService.NewUpdate(repo.userRepository, repo.userHistoryRepository, lib.redis),
			userService.NewDelete(repo.userRepository, lib.redis),
		),
		authService: authService.NewAuthService(
			repo.userRepository,
			repo.userHistoryRepository,
			lib.redis,
			lib.bcrypt,
		),
		provinceService: provinceService.NewProvinceService(
			provinceService.NewGet(repo.provinceRepository, lib.redis),
		),
		districtService: districtService.NewDistrictService(
			districtService.NewGet(repo.districtRepository, lib.redis),
		),
		subDistrictService: subDistrictService.NewSubDistrictService(
			subDistrictService.NewGet(repo.subDistrictRepository, lib.redis),
		),
	}
}
