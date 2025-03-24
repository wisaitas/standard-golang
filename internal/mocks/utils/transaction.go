package mock_utils

import (
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockTransactionUtil struct {
	mock.Mock
}

func NewMockTransactionUtil() *MockTransactionUtil {
	return &MockTransactionUtil{}
}

func (r *MockTransactionUtil) ExecuteInTransaction(fn func(tx *gorm.DB) error) error {
	args := r.Called(fn)
	return args.Error(0)
}

func (r *MockTransactionUtil) GetTransaction() *gorm.DB {
	args := r.Called()
	return args.Get(0).(*gorm.DB)
}

func (r *MockTransactionUtil) Begin() error {
	args := r.Called()
	return args.Error(0)
}

func (r *MockTransactionUtil) Commit() error {
	args := r.Called()
	return args.Error(0)
}

func (r *MockTransactionUtil) Rollback() error {
	args := r.Called()
	return args.Error(0)
}
