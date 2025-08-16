package initial

import (
	repositoryPkg "github.com/wisaitas/share-pkg/db/repository"
	"github.com/wisaitas/standard-golang/internal/standard-service/entity"
	repositoryInternal "github.com/wisaitas/standard-golang/internal/standard-service/repository"
)

type repository struct {
	userRepository        repositoryInternal.UserRepository
	provinceRepository    repositoryInternal.ProvinceRepository
	districtRepository    repositoryInternal.DistrictRepository
	subDistrictRepository repositoryInternal.SubDistrictRepository
}

func newRepository(clientConfig *clientConfig) *repository {
	return &repository{
		userRepository: repositoryInternal.NewUserRepository(
			clientConfig.DB,
			repositoryPkg.NewBaseRepository[entity.User](clientConfig.DB),
		),
		provinceRepository: repositoryInternal.NewProvinceRepository(
			clientConfig.DB,
			repositoryPkg.NewBaseRepository[entity.Province](clientConfig.DB),
		),
		districtRepository: repositoryInternal.NewDistrictRepository(
			clientConfig.DB,
			repositoryPkg.NewBaseRepository[entity.District](clientConfig.DB),
		),
		subDistrictRepository: repositoryInternal.NewSubDistrictRepository(
			clientConfig.DB,
			repositoryPkg.NewBaseRepository[entity.SubDistrict](clientConfig.DB),
		),
	}
}
