package initial

import (
	"github.com/wisaitas/standard-golang/internal/repositories"
	"gorm.io/gorm"
)

func initializeRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		UserRepository: repositories.NewUserRepository(db),
	}
}

type Repositories struct {
	UserRepository repositories.UserRepository
}
