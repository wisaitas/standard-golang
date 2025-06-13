package repository

import (
	"github.com/stretchr/testify/mock"
	"github.com/wisaitas/standard-golang/internal/standard-service/entity"
	"github.com/wisaitas/standard-golang/pkg"
	"gorm.io/gorm"
)

type MockUserRepository struct {
	mock.Mock
}

func (r *MockUserRepository) GetAll(items *[]entity.User, pagination *pkg.PaginationQuery, condition *pkg.Condition, relations *[]pkg.Relation) error {
	args := r.Called(items, pagination, condition, relations)
	return args.Error(0)
}

func (r *MockUserRepository) GetBy(item *entity.User, condition *pkg.Condition, relations *[]pkg.Relation) error {
	args := r.Called(item, condition, relations)
	return args.Error(0)
}

func (r *MockUserRepository) Create(item *entity.User) error {
	args := r.Called(item)
	return args.Error(0)
}

func (r *MockUserRepository) CreateMany(items *[]entity.User) error {
	args := r.Called(items)
	return args.Error(0)
}

func (r *MockUserRepository) Update(item *entity.User) error {
	args := r.Called(item)
	return args.Error(0)
}

func (r *MockUserRepository) UpdateMany(items *[]entity.User) error {
	args := r.Called(items)
	return args.Error(0)
}

func (r *MockUserRepository) Save(item *entity.User) error {
	args := r.Called(item)
	return args.Error(0)
}

func (r *MockUserRepository) SaveMany(items *[]entity.User) error {
	args := r.Called(items)
	return args.Error(0)
}

func (r *MockUserRepository) Delete(item *entity.User) error {
	args := r.Called(item)
	return args.Error(0)
}

func (r *MockUserRepository) DeleteMany(items *[]entity.User) error {
	args := r.Called(items)
	return args.Error(0)
}

func (r *MockUserRepository) WithTx(tx *gorm.DB) pkg.BaseRepository[entity.User] {
	args := r.Called(tx)
	return args.Get(0).(pkg.BaseRepository[entity.User])
}
