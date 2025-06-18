package port

import (
	"context"
	"go-gin-boilerplate/internal/domain"
)

type BarRepository interface {
	CreateBar(ctx context.Context, bar *domain.Bar) (*domain.Bar, error)
	GetBars(ctx context.Context) ([]*domain.Bar, error)
	GetBarByID(ctx context.Context, id string) (*domain.Bar, error)
	UpdateBar(ctx context.Context, id string, bar *domain.Bar) (*domain.Bar, error)
	DeleteBar(ctx context.Context, id string) error
}

type BarService interface {
	CreateBar(ctx context.Context, bar *domain.Bar) (*domain.Bar, error)
	GetBars(ctx context.Context) ([]*domain.Bar, error)
	GetBarByID(ctx context.Context, id string) (*domain.Bar, error)
	UpdateBar(ctx context.Context, id string, bar *domain.Bar) (*domain.Bar, error)
	DeleteBar(ctx context.Context, id string) error
}
