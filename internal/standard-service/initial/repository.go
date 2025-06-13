package initial

import (
	"github.com/wisaitas/standard-golang/internal/standard-service/entity"
	repositoryInternal "github.com/wisaitas/standard-golang/internal/standard-service/repository"
	"github.com/wisaitas/standard-golang/pkg"
)

type repository struct {
	userRepository        repositoryInternal.UserRepository
	userHistoryRepository repositoryInternal.UserHistoryRepository
	provinceRepository    repositoryInternal.ProvinceRepository
	districtRepository    repositoryInternal.DistrictRepository
	subDistrictRepository repositoryInternal.SubDistrictRepository
}

func newRepository(clientConfig *clientConfig) *repository {
	return &repository{
		userRepository: repositoryInternal.NewUserRepository(
			clientConfig.DB,
			pkg.NewBaseRepository[entity.User](clientConfig.DB),
		),
		userHistoryRepository: repositoryInternal.NewUserHistoryRepository(
			clientConfig.DB,
			pkg.NewBaseRepository[entity.UserHistory](clientConfig.DB),
		),
		provinceRepository: repositoryInternal.NewProvinceRepository(
			clientConfig.DB,
			pkg.NewBaseRepository[entity.Province](clientConfig.DB),
		),
		districtRepository: repositoryInternal.NewDistrictRepository(
			clientConfig.DB,
			pkg.NewBaseRepository[entity.District](clientConfig.DB),
		),
		subDistrictRepository: repositoryInternal.NewSubDistrictRepository(
			clientConfig.DB,
			pkg.NewBaseRepository[entity.SubDistrict](clientConfig.DB),
		),
	}
}
