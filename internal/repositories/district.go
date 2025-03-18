package repositories

import (
	"github.com/wisaitas/standard-golang/internal/models"
	"github.com/wisaitas/standard-golang/pkg"
	"gorm.io/gorm"
)

type DistrictRepository interface {
	pkg.BaseRepository[models.District]
}

type districtRepository struct {
	pkg.BaseRepository[models.District]
	db *gorm.DB
}

func NewDistrictRepository(db *gorm.DB, baseRepository pkg.BaseRepository[models.District]) DistrictRepository {
	return &districtRepository{
		BaseRepository: baseRepository,
		db:             db,
	}
}
