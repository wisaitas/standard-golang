package repositories

import (
	"github.com/wisaitas/standard-golang/internal/models"
	"github.com/wisaitas/standard-golang/pkg"

	"gorm.io/gorm"
)

type UserRepository interface {
	pkg.BaseRepository[models.User]
}

type userRepository struct {
	pkg.BaseRepository[models.User]
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB, baseRepository pkg.BaseRepository[models.User]) UserRepository {
	return &userRepository{
		BaseRepository: baseRepository,
		db:             db,
	}
}
