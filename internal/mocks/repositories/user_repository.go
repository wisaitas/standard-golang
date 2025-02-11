package mock_repositories

import (
	"github.com/wisaitas/standard-golang/internal/dtos/request"

	"github.com/wisaitas/standard-golang/internal/models"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

// MockUserRepository is a mock implementation of UserRepository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetAll(items *[]models.User, pagination *request.PaginationParam) error {
	args := m.Called(items, pagination)
	return args.Error(0)
}

func (m *MockUserRepository) GetById(id uuid.UUID, item *models.User) error {
	args := m.Called(id, item)
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

func (m *MockUserRepository) Save(item *models.User) error {
	args := m.Called(item)
	return args.Error(0)
}

func (m *MockUserRepository) Delete(item *models.User) error {
	args := m.Called(item)
	return args.Error(0)
}

func (m *MockUserRepository) GetBy(field string, value string, item *models.User) error {
	args := m.Called(field, value, item)
	return args.Error(0)
}
