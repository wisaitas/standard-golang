package initial

import (
	"github.com/wisaitas/standard-golang/internal/models"
	"github.com/wisaitas/standard-golang/internal/repositories"
	"gorm.io/gorm"
)

type Repositories struct {
	UserRepository        repositories.UserRepository
	ProvinceRepository    repositories.ProvinceRepository
	DistrictRepository    repositories.DistrictRepository
	SubDistrictRepository repositories.SubDistrictRepository
}

func initializeRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		UserRepository:        repositories.NewUserRepository(db, repositories.NewBaseRepository[models.User](db)),
		ProvinceRepository:    repositories.NewProvinceRepository(db, repositories.NewBaseRepository[models.Province](db)),
		DistrictRepository:    repositories.NewDistrictRepository(db, repositories.NewBaseRepository[models.District](db)),
		SubDistrictRepository: repositories.NewSubDistrictRepository(db, repositories.NewBaseRepository[models.SubDistrict](db)),
	}
}
