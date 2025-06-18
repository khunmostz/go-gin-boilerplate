package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
)

// MockBaseRepository implements db.BaseRepository interface for testing
type MockBaseRepository struct {
	mock.Mock
}

func (m *MockBaseRepository) Create(ctx context.Context, collection string, data interface{}) error {
	args := m.Called(ctx, collection, data)
	return args.Error(0)
}

func (m *MockBaseRepository) GetAll(ctx context.Context, collection string, result interface{}) error {
	args := m.Called(ctx, collection, result)
	return args.Error(0)
}

func (m *MockBaseRepository) GetById(ctx context.Context, collection string, id string, result interface{}) error {
	args := m.Called(ctx, collection, id, result)
	return args.Error(0)
}

func (m *MockBaseRepository) GetByField(ctx context.Context, collection string, field string, value interface{}, result interface{}) error {
	args := m.Called(ctx, collection, field, value, result)
	return args.Error(0)
}

func (m *MockBaseRepository) UpdateById(ctx context.Context, collection string, id string, data interface{}) error {
	args := m.Called(ctx, collection, id, data)
	return args.Error(0)
}

func (m *MockBaseRepository) DeleteById(ctx context.Context, collection string, id string) error {
	args := m.Called(ctx, collection, id)
	return args.Error(0)
}
