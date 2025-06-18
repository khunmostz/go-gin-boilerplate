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

func (m *MockFooRepository) CreateFoo(ctx context.Context, foo *domain.Foo) (*domain.Foo, error) {
	args := m.Called(ctx, foo)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Foo), args.Error(1)
}

// MockFooService implements port.FooService interface for testing
type MockFooService struct {
	mock.Mock
}

func (m *MockFooService) CreateFoo(ctx context.Context, foo *domain.Foo) (*domain.Foo, error) {
	args := m.Called(ctx, foo)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Foo), args.Error(1)
}
