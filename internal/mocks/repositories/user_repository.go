package mock_repositories

import (
	"github.com/wisaitas/standard-golang/internal/dtos/queries"
	"github.com/wisaitas/standard-golang/internal/repositories"
	"gorm.io/gorm"

	"github.com/wisaitas/standard-golang/internal/models"

	"github.com/stretchr/testify/mock"
)

// MockUserRepository is a mock implementation of UserRepository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) WithTx(tx *gorm.DB) repositories.BaseRepository[models.User] {
	args := m.Called(tx)
	return args.Get(0).(repositories.BaseRepository[models.User])
}

func (m *MockUserRepository) GetAll(items *[]models.User, pagination *queries.PaginationQuery, condition interface{}, relations ...string) error {
	args := m.Called(items, pagination, condition, relations)
	return args.Error(0)
}

func (m *MockUserRepository) GetBy(condition interface{}, item *models.User, relations ...string) error {
	args := m.Called(condition, item, relations)
	return args.Error(0)
}

func (m *MockUserRepository) Create(item *models.User) error {
	args := m.Called(item)
	return args.Error(0)
}

func (m *MockUserRepository) CreateMany(items *[]models.User) error {
	args := m.Called(items)
	return args.Error(0)
}

func (m *MockUserRepository) Updates(item *models.User) error {
	args := m.Called(item)
	return args.Error(0)
}

func (m *MockUserRepository) UpdateMany(items *[]models.User) error {
	args := m.Called(items)
	return args.Error(0)
}

func (m *MockUserRepository) Save(item *models.User) error {
	args := m.Called(item)
	return args.Error(0)
}

func (m *MockUserRepository) SaveMany(items *[]models.User) error {
	args := m.Called(items)
	return args.Error(0)
}

func (m *MockUserRepository) Delete(item *models.User) error {
	args := m.Called(item)
	return args.Error(0)
}
