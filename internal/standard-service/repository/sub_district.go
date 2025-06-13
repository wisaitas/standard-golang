package repository

import (
	"github.com/wisaitas/standard-golang/internal/standard-service/entity"
	"github.com/wisaitas/standard-golang/pkg"
	"gorm.io/gorm"
)

type SubDistrictRepository interface {
	pkg.BaseRepository[entity.SubDistrict]
}

type subDistrictRepository struct {
	pkg.BaseRepository[entity.SubDistrict]
	db *gorm.DB
}

func NewSubDistrictRepository(
	db *gorm.DB,
	baseRepository pkg.BaseRepository[entity.SubDistrict],
) SubDistrictRepository {
	return &subDistrictRepository{
		BaseRepository: baseRepository,
		db:             db,
	}
}
