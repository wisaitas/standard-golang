package initial

import (
	"github.com/wisaitas/standard-golang/internal/models"
	"github.com/wisaitas/standard-golang/internal/repositories"
	"github.com/wisaitas/standard-golang/pkg"
	"gorm.io/gorm"
)

type Repository struct {
	UserRepository        repositories.UserRepository
	UserHistoryRepository repositories.UserHistoryRepository
	ProvinceRepository    repositories.ProvinceRepository
	DistrictRepository    repositories.DistrictRepository
	SubDistrictRepository repositories.SubDistrictRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		UserRepository:        repositories.NewUserRepository(db, pkg.NewBaseRepository[models.User](db)),
		UserHistoryRepository: repositories.NewUserHistoryRepository(db, pkg.NewBaseRepository[models.UserHistory](db)),
		ProvinceRepository:    repositories.NewProvinceRepository(db, pkg.NewBaseRepository[models.Province](db)),
		DistrictRepository:    repositories.NewDistrictRepository(db, pkg.NewBaseRepository[models.District](db)),
		SubDistrictRepository: repositories.NewSubDistrictRepository(db, pkg.NewBaseRepository[models.SubDistrict](db)),
	}
}
