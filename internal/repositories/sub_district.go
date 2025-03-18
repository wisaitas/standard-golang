package repositories

import (
	"github.com/wisaitas/standard-golang/internal/models"
	"github.com/wisaitas/standard-golang/pkg"
	"gorm.io/gorm"
)

type SubDistrictRepository interface {
	pkg.BaseRepository[models.SubDistrict]
}

type subDistrictRepository struct {
	pkg.BaseRepository[models.SubDistrict]
	db *gorm.DB
}

func NewSubDistrictRepository(db *gorm.DB, baseRepository pkg.BaseRepository[models.SubDistrict]) SubDistrictRepository {
	return &subDistrictRepository{
		BaseRepository: baseRepository,
		db:             db,
	}
}
