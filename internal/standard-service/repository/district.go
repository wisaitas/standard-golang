package repository

import (
	"github.com/wisaitas/share-pkg/db/repository"
	"github.com/wisaitas/standard-golang/internal/standard-service/entity"
	"gorm.io/gorm"
)

type DistrictRepository interface {
	repository.BaseRepository[entity.District]
}

type districtRepository struct {
	repository.BaseRepository[entity.District]
	db *gorm.DB
}

func NewDistrictRepository(
	db *gorm.DB,
	baseRepository repository.BaseRepository[entity.District],
) DistrictRepository {
	return &districtRepository{
		BaseRepository: baseRepository,
		db:             db,
	}
}
