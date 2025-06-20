package mocks

import (
	"context"
	"go-gin-boilerplate/internal/domain"

	"github.com/stretchr/testify/mock"
)

// MockFooRepository implements port.FooRepository interface for testing
type MockFooRepository struct {
	mock.Mock
}

func (m *MockFooRepository) Create(ctx context.Context, foo *domain.Foo) (*domain.Foo, error) {
	args := m.Called(ctx, foo)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Foo), args.Error(1)
}

func (m *MockFooRepository) GetAll(ctx context.Context) ([]*domain.Foo, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Foo), args.Error(1)
}

func (m *MockFooRepository) GetByID(ctx context.Context, id string) (*domain.Foo, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Foo), args.Error(1)
}

func (m *MockFooRepository) GetByName(ctx context.Context, name string) (*domain.Foo, error) {
	args := m.Called(ctx, name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Foo), args.Error(1)
}

func (m *MockFooRepository) UpdateById(ctx context.Context, id string, update map[string]any) (*domain.Foo, error) {
	args := m.Called(ctx, id, update)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Foo), args.Error(1)
}

func (m *MockFooRepository) DeleteById(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// MockFooService implements port.FooService interface for testing
type MockFooService struct {
	mock.Mock
}

func (m *MockFooService) Create(ctx context.Context, foo *domain.Foo) (*domain.Foo, error) {
	args := m.Called(ctx, foo)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Foo), args.Error(1)
}

func (m *MockFooService) GetAll(ctx context.Context) ([]*domain.Foo, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Foo), args.Error(1)
}

func (m *MockFooService) GetByID(ctx context.Context, id string) (*domain.Foo, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Foo), args.Error(1)
}

func (m *MockFooService) GetByName(ctx context.Context, name string) (*domain.Foo, error) {
	args := m.Called(ctx, name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Foo), args.Error(1)
}

func (m *MockFooService) UpdateById(ctx context.Context, id string, update map[string]any) (*domain.Foo, error) {
	args := m.Called(ctx, id, update)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Foo), args.Error(1)
}

func (m *MockFooService) DeleteById(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
