package initial

import (
	"github.com/wisaitas/standard-golang/internal/models"
	"github.com/wisaitas/standard-golang/internal/repositories"
	"gorm.io/gorm"
)

func initializeRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		UserRepository: repositories.NewUserRepository(db, repositories.NewBaseRepository[models.User](db)),
	}
}

type Repositories struct {
	UserRepository repositories.UserRepository
}
