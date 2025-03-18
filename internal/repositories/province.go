package repositories

import (
	"github.com/wisaitas/standard-golang/internal/models"
	"github.com/wisaitas/standard-golang/pkg"
	"gorm.io/gorm"
)

type ProvinceRepository interface {
	pkg.BaseRepository[models.Province]
}

type provinceRepository struct {
	pkg.BaseRepository[models.Province]
	db *gorm.DB
}

func NewProvinceRepository(db *gorm.DB, baseRepository pkg.BaseRepository[models.Province]) ProvinceRepository {
	return &provinceRepository{
		BaseRepository: baseRepository,
		db:             db,
	}
}
