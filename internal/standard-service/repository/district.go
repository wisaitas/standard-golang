package repository

import (
	"github.com/wisaitas/standard-golang/internal/standard-service/entity"
	"github.com/wisaitas/standard-golang/pkg"
	"gorm.io/gorm"
)

type DistrictRepository interface {
	pkg.BaseRepository[entity.District]
}

type districtRepository struct {
	pkg.BaseRepository[entity.District]
	db *gorm.DB
}

func NewDistrictRepository(
	db *gorm.DB,
	baseRepository pkg.BaseRepository[entity.District],
) DistrictRepository {
	return &districtRepository{
		BaseRepository: baseRepository,
		db:             db,
	}
}
