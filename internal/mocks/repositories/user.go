package mock_repositories

import (
	"gorm.io/gorm"

	"github.com/wisaitas/standard-golang/internal/models"
	"github.com/wisaitas/standard-golang/pkg"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (r *MockUserRepository) GetAll(items *[]models.User, pagination *pkg.PaginationQuery, condition interface{}, relations ...string) error {
	args := r.Called(items, pagination, condition, relations)
	return args.Error(0)
}

func (r *MockUserRepository) GetBy(condition interface{}, item *models.User, relations ...string) error {
	args := r.Called(condition, item, relations)
	return args.Error(0)
}

func (r *MockUserRepository) Create(item *models.User) error {
	args := r.Called(item)
	return args.Error(0)
}

func (r *MockUserRepository) CreateMany(items *[]models.User) error {
	args := r.Called(items)
	return args.Error(0)
}

func (r *MockUserRepository) Update(item *models.User) error {
	args := r.Called(item)
	return args.Error(0)
}

func (r *MockUserRepository) UpdateMany(items *[]models.User) error {
	args := r.Called(items)
	return args.Error(0)
}

func (r *MockUserRepository) Save(item *models.User) error {
	args := r.Called(item)
	return args.Error(0)
}

func (r *MockUserRepository) SaveMany(items *[]models.User) error {
	args := r.Called(items)
	return args.Error(0)
}

func (r *MockUserRepository) Delete(item *models.User) error {
	args := r.Called(item)
	return args.Error(0)
}

func (r *MockUserRepository) DeleteMany(items *[]models.User) error {
	args := r.Called(items)
	return args.Error(0)
}

func (r *MockUserRepository) WithTx(tx *gorm.DB) pkg.BaseRepository[models.User] {
	args := r.Called(tx)
	return args.Get(0).(pkg.BaseRepository[models.User])
}
