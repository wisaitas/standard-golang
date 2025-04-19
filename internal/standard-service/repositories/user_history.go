package repositories

import (
	"github.com/wisaitas/standard-golang/internal/standard-service/models"
	"github.com/wisaitas/standard-golang/pkg"
	"gorm.io/gorm"
)

type UserHistoryRepository interface {
	pkg.BaseRepository[models.UserHistory]
}

type userHistoryRepository struct {
	pkg.BaseRepository[models.UserHistory]
	db *gorm.DB
}

func NewUserHistoryRepository(db *gorm.DB, baseRepository pkg.BaseRepository[models.UserHistory]) UserHistoryRepository {
	return &userHistoryRepository{
		BaseRepository: baseRepository,
		db:             db,
	}
}
