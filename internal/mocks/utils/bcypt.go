package mock_utils

import "github.com/stretchr/testify/mock"

type MockBcrypt struct {
	mock.Mock
}

func NewMockBcrypt() *MockBcrypt {
	return &MockBcrypt{}
}

func (r *MockBcrypt) GenerateFromPassword(password string, cost int) ([]byte, error) {
	args := r.Called(password, cost)
	return args.Get(0).([]byte), args.Error(1)
}

func (r *MockBcrypt) CompareHashAndPassword(hashedPassword, password []byte) error {
	args := r.Called(hashedPassword, password)
	return args.Error(0)
}
