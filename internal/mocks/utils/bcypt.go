package mock_utils

import "github.com/stretchr/testify/mock"

type MockBcrypt struct {
	mock.Mock
}

func (m *MockBcrypt) GenerateFromPassword(password string, cost int) ([]byte, error) {
	args := m.Called(password, cost)
	return args.Get(0).([]byte), args.Error(1)
}

func (m *MockBcrypt) CompareHashAndPassword(hashedPassword, password []byte) error {
	args := m.Called(hashedPassword, password)
	return args.Error(0)
}
