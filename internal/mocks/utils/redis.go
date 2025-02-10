package mock_utils

import (
	"context"
	"time"

	"github.com/stretchr/testify/mock"
)

type MockRedis struct {
	mock.Mock
}

func (m *MockRedis) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	args := m.Called(ctx, key, value, expiration)
	return args.Error(0)
}

func (m *MockRedis) Get(ctx context.Context, key string) (string, error) {
	args := m.Called(ctx, key)
	return args.String(0), args.Error(1)
}

func (m *MockRedis) Del(ctx context.Context, keys ...string) error {
	args := m.Called(ctx, keys)
	return args.Error(0)
}

func (m *MockRedis) Exists(ctx context.Context, keys ...string) (bool, error) {
	args := m.Called(ctx, keys)
	return args.Bool(0), args.Error(1)
}
