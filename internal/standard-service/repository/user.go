package repository

import (
	"github.com/wisaitas/standard-golang/internal/standard-service/entity"
	"github.com/wisaitas/standard-golang/pkg"

	"gorm.io/gorm"
)

type UserRepository interface {
	pkg.BaseRepository[entity.User]
}

type userRepository struct {
	pkg.BaseRepository[entity.User]
	db *gorm.DB
}

func NewUserRepository(
	db *gorm.DB,
	baseRepository pkg.BaseRepository[entity.User],
) UserRepository {
	return &userRepository{
		BaseRepository: baseRepository,
		db:             db,
	}
}
