package initial

import (
	"github.com/wisaitas/standard-golang/internal/models"
	"github.com/wisaitas/standard-golang/internal/repositories"
	"github.com/wisaitas/standard-golang/pkg"
)

type repository struct {
	userRepository        repositories.UserRepository
	userHistoryRepository repositories.UserHistoryRepository
	provinceRepository    repositories.ProvinceRepository
	districtRepository    repositories.DistrictRepository
	subDistrictRepository repositories.SubDistrictRepository
}

func newRepository(config *config) *repository {
	return &repository{
		userRepository:        repositories.NewUserRepository(config.DB, pkg.NewBaseRepository[models.User](config.DB)),
		userHistoryRepository: repositories.NewUserHistoryRepository(config.DB, pkg.NewBaseRepository[models.UserHistory](config.DB)),
		provinceRepository:    repositories.NewProvinceRepository(config.DB, pkg.NewBaseRepository[models.Province](config.DB)),
		districtRepository:    repositories.NewDistrictRepository(config.DB, pkg.NewBaseRepository[models.District](config.DB)),
		subDistrictRepository: repositories.NewSubDistrictRepository(config.DB, pkg.NewBaseRepository[models.SubDistrict](config.DB)),
	}
}
