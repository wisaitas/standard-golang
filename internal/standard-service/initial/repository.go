package initial

import (
	"github.com/wisaitas/standard-golang/internal/standard-service/models"
	"github.com/wisaitas/standard-golang/internal/standard-service/repositories"
	"github.com/wisaitas/standard-golang/pkg"
)

type repository struct {
	userRepository        repositories.UserRepository
	userHistoryRepository repositories.UserHistoryRepository
	provinceRepository    repositories.ProvinceRepository
	districtRepository    repositories.DistrictRepository
	subDistrictRepository repositories.SubDistrictRepository
}

func newRepository(clientConfig *clientConfig) *repository {
	return &repository{
		userRepository:        repositories.NewUserRepository(clientConfig.DB, pkg.NewBaseRepository[models.User](clientConfig.DB)),
		userHistoryRepository: repositories.NewUserHistoryRepository(clientConfig.DB, pkg.NewBaseRepository[models.UserHistory](clientConfig.DB)),
		provinceRepository:    repositories.NewProvinceRepository(clientConfig.DB, pkg.NewBaseRepository[models.Province](clientConfig.DB)),
		districtRepository:    repositories.NewDistrictRepository(clientConfig.DB, pkg.NewBaseRepository[models.District](clientConfig.DB)),
		subDistrictRepository: repositories.NewSubDistrictRepository(clientConfig.DB, pkg.NewBaseRepository[models.SubDistrict](clientConfig.DB)),
	}
}
