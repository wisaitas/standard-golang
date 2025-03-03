package repositories

import (
	"github.com/wisaitas/standard-golang/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	BaseRepository[models.User]
}

type userRepository struct {
	BaseRepository[models.User]
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB, baseRepository BaseRepository[models.User]) UserRepository {
	return &userRepository{
		BaseRepository: baseRepository,
		db:             db,
	}
}
