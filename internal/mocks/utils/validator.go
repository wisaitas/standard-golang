package mock_utils

import "github.com/stretchr/testify/mock"

type MockValidate struct {
	mock.Mock
}

func NewMockValidate() *MockValidate {
	return &MockValidate{}
}

func (r *MockValidate) Validate(s interface{}) error {
	args := r.Called(s)
	return args.Error(0)
}
