package pkg

import (
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockTransactionManager struct {
	mock.Mock
}

func NewMockTransactionManager() *MockTransactionManager {
	return &MockTransactionManager{}
}

func (r *MockTransactionManager) ExecuteInTransaction(fn func(tx *gorm.DB) error) error {
	args := r.Called(fn)
	return args.Error(0)
}

func (r *MockTransactionManager) GetTransaction() *gorm.DB {
	args := r.Called()
	return args.Get(0).(*gorm.DB)
}

func (r *MockTransactionManager) Begin() error {
	args := r.Called()
	return args.Error(0)
}

func (r *MockTransactionManager) Commit() error {
	args := r.Called()
	return args.Error(0)
}

func (r *MockTransactionManager) Rollback() error {
	args := r.Called()
	return args.Error(0)
}
