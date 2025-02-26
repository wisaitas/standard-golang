package repositories

import (
	"github.com/wisaitas/standard-golang/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	BaseRepository[models.User]
	GetUsersPreloadAddresses(users *[]models.User) error
}

type userRepository struct {
	BaseRepository[models.User]
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		BaseRepository: NewBaseRepository[models.User](db),
		db:             db,
	}
}

func (r *userRepository) GetUsersPreloadAddresses(users *[]models.User) error {
	return r.db.Preload("Addresses").Find(&users).Error
}
