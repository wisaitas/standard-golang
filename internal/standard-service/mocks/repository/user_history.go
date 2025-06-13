package repository

import (
	"github.com/stretchr/testify/mock"
	"github.com/wisaitas/standard-golang/internal/standard-service/entity"
	"github.com/wisaitas/standard-golang/pkg"
	"gorm.io/gorm"
)

type MockUserHistoryRepository struct {
	mock.Mock
}

func (r *MockUserHistoryRepository) GetAll(items *[]entity.UserHistory, pagination *pkg.PaginationQuery, condition *pkg.Condition, relations *[]pkg.Relation) error {
	args := r.Called(items, pagination, condition, relations)
	return args.Error(0)
}

func (r *MockUserHistoryRepository) GetBy(item *entity.UserHistory, condition *pkg.Condition, relations *[]pkg.Relation) error {
	args := r.Called(item, condition, relations)
	return args.Error(0)
}

func (r *MockUserHistoryRepository) Create(item *entity.UserHistory) error {
	args := r.Called(item)
	return args.Error(0)
}

func (r *MockUserHistoryRepository) CreateMany(items *[]entity.UserHistory) error {
	args := r.Called(items)
	return args.Error(0)
}

func (r *MockUserHistoryRepository) Update(item *entity.UserHistory) error {
	args := r.Called(item)
	return args.Error(0)
}

func (r *MockUserHistoryRepository) UpdateMany(items *[]entity.UserHistory) error {
	args := r.Called(items)
	return args.Error(0)
}

func (r *MockUserHistoryRepository) Save(item *entity.UserHistory) error {
	args := r.Called(item)
	return args.Error(0)
}

func (r *MockUserHistoryRepository) SaveMany(items *[]entity.UserHistory) error {
	args := r.Called(items)
	return args.Error(0)
}

func (r *MockUserHistoryRepository) Delete(item *entity.UserHistory) error {
	args := r.Called(item)
	return args.Error(0)
}

func (r *MockUserHistoryRepository) DeleteMany(items *[]entity.UserHistory) error {
	args := r.Called(items)
	return args.Error(0)
}

func (r *MockUserHistoryRepository) WithTxManager(tm *pkg.TxManager) pkg.BaseRepository[entity.UserHistory] {
	args := r.Called(tm)
	return args.Get(0).(pkg.BaseRepository[entity.UserHistory])
}

func (r *MockUserHistoryRepository) GetDB() *gorm.DB {
	args := r.Called()
	return args.Get(0).(*gorm.DB)
}
