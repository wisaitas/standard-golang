package repository

import (
	"github.com/wisaitas/share-pkg/db/repository"
	"github.com/wisaitas/standard-golang/internal/standard-service/entity"
	"gorm.io/gorm"
)

type SubDistrictRepository interface {
	repository.BaseRepository[entity.SubDistrict]
}

type subDistrictRepository struct {
	repository.BaseRepository[entity.SubDistrict]
	db *gorm.DB
}

func NewSubDistrictRepository(
	db *gorm.DB,
	baseRepository repository.BaseRepository[entity.SubDistrict],
) SubDistrictRepository {
	return &subDistrictRepository{
		BaseRepository: baseRepository,
		db:             db,
	}
}
