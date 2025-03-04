package repositories

import (
	"github.com/wisaitas/standard-golang/internal/models"
	"gorm.io/gorm"
)

type DistrictRepository interface {
	BaseRepository[models.District]
}

type districtRepository struct {
	BaseRepository[models.District]
	db *gorm.DB
}

func NewDistrictRepository(db *gorm.DB, baseRepository BaseRepository[models.District]) DistrictRepository {
	return &districtRepository{
		BaseRepository: baseRepository,
		db:             db,
	}
}
