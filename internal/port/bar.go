package port

import (
	"context"
	"go-gin-boilerplate/internal/domain"
)

type BarRepository interface {
	Create(ctx context.Context, bar *domain.Bar) (*domain.Bar, error)
	GetAll(ctx context.Context) ([]*domain.Bar, error)
	GetByID(ctx context.Context, id string) (*domain.Bar, error)
	GetByName(ctx context.Context, name string) (*domain.Bar, error)
	UpdateById(ctx context.Context, id string, update map[string]any) (*domain.Bar, error)
	DeleteById(ctx context.Context, id string) error
}

type BarService interface {
	Create(ctx context.Context, bar *domain.Bar) (*domain.Bar, error)
	GetAll(ctx context.Context) ([]*domain.Bar, error)
	GetByID(ctx context.Context, id string) (*domain.Bar, error)
	GetByName(ctx context.Context, name string) (*domain.Bar, error)
	UpdateById(ctx context.Context, id string, update map[string]any) (*domain.Bar, error)
	DeleteById(ctx context.Context, id string) error
}
