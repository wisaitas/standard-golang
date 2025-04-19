package mock_utils

import (
	"context"
	"time"

	"github.com/stretchr/testify/mock"
)

type MockRedis struct {
	mock.Mock
}

func (r *MockRedis) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	args := r.Called(ctx, key, value, expiration)
	return args.Error(0)
}

func (r *MockRedis) Get(ctx context.Context, key string) (string, error) {
	args := r.Called(ctx, key)
	return args.String(0), args.Error(1)
}

func (r *MockRedis) Del(ctx context.Context, keys ...string) error {
	args := r.Called(ctx, keys)
	return args.Error(0)
}

func (r *MockRedis) Exists(ctx context.Context, keys ...string) (bool, error) {
	args := r.Called(ctx, keys)
	return args.Bool(0), args.Error(1)
}
