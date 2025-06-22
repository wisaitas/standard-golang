package repository

import (
	"github.com/wisaitas/share-pkg/db/repository"
	"github.com/wisaitas/standard-golang/internal/standard-service/entity"
	"gorm.io/gorm"
)

type ProvinceRepository interface {
	repository.BaseRepository[entity.Province]
}

type provinceRepository struct {
	repository.BaseRepository[entity.Province]
	db *gorm.DB
}

func NewProvinceRepository(
	db *gorm.DB,
	baseRepository repository.BaseRepository[entity.Province],
) ProvinceRepository {
	return &provinceRepository{
		BaseRepository: baseRepository,
		db:             db,
	}
}
