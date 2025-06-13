package repository

import (
	"github.com/wisaitas/standard-golang/internal/standard-service/entity"
	"github.com/wisaitas/standard-golang/pkg"
	"gorm.io/gorm"
)

type UserHistoryRepository interface {
	pkg.BaseRepository[entity.UserHistory]
}

type userHistoryRepository struct {
	pkg.BaseRepository[entity.UserHistory]
	db *gorm.DB
}

func NewUserHistoryRepository(
	db *gorm.DB,
	baseRepository pkg.BaseRepository[entity.UserHistory],
) UserHistoryRepository {
	return &userHistoryRepository{
		BaseRepository: baseRepository,
		db:             db,
	}
}
