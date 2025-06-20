package mocks

import (
	"context"
	"go-gin-boilerplate/internal/domain"

	"github.com/stretchr/testify/mock"
)

type MockBarRepository struct {
	mock.Mock
}

func (m *MockBarRepository) Create(ctx context.Context, bar *domain.Bar) (*domain.Bar, error) {
	args := m.Called(ctx, bar)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Bar), args.Error(1)
}

func (m *MockBarRepository) GetAll(ctx context.Context) ([]*domain.Bar, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Bar), args.Error(1)
}

func (m *MockBarRepository) GetByID(ctx context.Context, id string) (*domain.Bar, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Bar), args.Error(1)
}

func (m *MockBarRepository) GetByName(ctx context.Context, name string) (*domain.Bar, error) {
	args := m.Called(ctx, name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Bar), args.Error(1)
}

func (m *MockBarRepository) UpdateById(ctx context.Context, id string, update map[string]any) (*domain.Bar, error) {
	args := m.Called(ctx, id, update)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Bar), args.Error(1)
}

func (m *MockBarRepository) DeleteById(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// MockBarService implements port.BarService interface for testing
type MockBarService struct {
	mock.Mock
}

func (m *MockBarService) Create(ctx context.Context, bar *domain.Bar) (*domain.Bar, error) {
	args := m.Called(ctx, bar)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Bar), args.Error(1)
}

func (m *MockBarService) GetAll(ctx context.Context) ([]*domain.Bar, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Bar), args.Error(1)
}

func (m *MockBarService) GetByID(ctx context.Context, id string) (*domain.Bar, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Bar), args.Error(1)
}

func (m *MockBarService) GetByName(ctx context.Context, name string) (*domain.Bar, error) {
	args := m.Called(ctx, name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Bar), args.Error(1)
}

func (m *MockBarService) UpdateById(ctx context.Context, id string, update map[string]any) (*domain.Bar, error) {
	args := m.Called(ctx, id, update)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Bar), args.Error(1)
}

func (m *MockBarService) DeleteById(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
