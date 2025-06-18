package mocks

import (
	"context"
	"go-gin-boilerplate/internal/domain"

	"github.com/stretchr/testify/mock"
)

type MockBarRepository struct {
	mock.Mock
}

func (m *MockBarRepository) CreateBar(ctx context.Context, bar *domain.Bar) (*domain.Bar, error) {
	args := m.Called(ctx, bar)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Bar), args.Error(1)
}
