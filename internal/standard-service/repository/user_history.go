package repository

import (
	"github.com/wisaitas/share-pkg/db/repository"
	"github.com/wisaitas/standard-golang/internal/standard-service/entity"
	"gorm.io/gorm"
)

type UserHistoryRepository interface {
	repository.BaseRepository[entity.UserHistory]
}

type userHistoryRepository struct {
	repository.BaseRepository[entity.UserHistory]
	db *gorm.DB
}

func NewUserHistoryRepository(
	db *gorm.DB,
	baseRepository repository.BaseRepository[entity.UserHistory],
) UserHistoryRepository {
	return &userHistoryRepository{
		BaseRepository: baseRepository,
		db:             db,
	}
}
