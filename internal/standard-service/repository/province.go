package repository

import (
	"github.com/wisaitas/standard-golang/internal/standard-service/entity"
	"github.com/wisaitas/standard-golang/pkg"
	"gorm.io/gorm"
)

type ProvinceRepository interface {
	pkg.BaseRepository[entity.Province]
}

type provinceRepository struct {
	pkg.BaseRepository[entity.Province]
	db *gorm.DB
}

func NewProvinceRepository(
	db *gorm.DB,
	baseRepository pkg.BaseRepository[entity.Province],
) ProvinceRepository {
	return &provinceRepository{
		BaseRepository: baseRepository,
		db:             db,
	}
}
