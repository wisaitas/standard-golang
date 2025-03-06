package repositories

import (
	"github.com/wisaitas/standard-golang/internal/models"
	"gorm.io/gorm"
)

type UserHistoryRepository interface {
	BaseRepository[models.UserHistory]
}

type userHistoryRepository struct {
	BaseRepository[models.UserHistory]
	db *gorm.DB
}

func NewUserHistoryRepository(db *gorm.DB, baseRepository BaseRepository[models.UserHistory]) UserHistoryRepository {
	return &userHistoryRepository{
		BaseRepository: baseRepository,
		db:             db,
	}
}
