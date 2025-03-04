package repositories

import (
	"github.com/wisaitas/standard-golang/internal/models"
	"gorm.io/gorm"
)

type ProvinceRepository interface {
	BaseRepository[models.Province]
}

type provinceRepository struct {
	BaseRepository[models.Province]
	db *gorm.DB
}

func NewProvinceRepository(db *gorm.DB, baseRepository BaseRepository[models.Province]) ProvinceRepository {
	return &provinceRepository{
		BaseRepository: baseRepository,
		db:             db,
	}
}
