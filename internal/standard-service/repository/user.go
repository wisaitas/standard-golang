package repository

import (
	"github.com/wisaitas/share-pkg/db/repository"
	"github.com/wisaitas/standard-golang/internal/standard-service/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	repository.BaseRepository[entity.User]
}

type userRepository struct {
	repository.BaseRepository[entity.User]
	db *gorm.DB
}

func NewUserRepository(
	db *gorm.DB,
	baseRepository repository.BaseRepository[entity.User],
) UserRepository {
	return &userRepository{
		BaseRepository: baseRepository,
		db:             db,
	}
}
