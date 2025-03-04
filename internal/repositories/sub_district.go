package repositories

import (
	"github.com/wisaitas/standard-golang/internal/models"
	"gorm.io/gorm"
)

type SubDistrictRepository interface {
	BaseRepository[models.SubDistrict]
}

type subDistrictRepository struct {
	BaseRepository[models.SubDistrict]
	db *gorm.DB
}

func NewSubDistrictRepository(db *gorm.DB, baseRepository BaseRepository[models.SubDistrict]) SubDistrictRepository {
	return &subDistrictRepository{
		BaseRepository: baseRepository,
		db:             db,
	}
}
