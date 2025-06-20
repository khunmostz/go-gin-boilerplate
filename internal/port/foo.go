package port

import (
	"context"
	"go-gin-boilerplate/internal/domain"
)

type FooRepository interface {
	Create(ctx context.Context, foo *domain.Foo) (*domain.Foo, error)
	GetAll(ctx context.Context) ([]*domain.Foo, error)
	GetByID(ctx context.Context, id string) (*domain.Foo, error)
	GetByName(ctx context.Context, name string) (*domain.Foo, error)
	UpdateById(ctx context.Context, id string, update map[string]any) (*domain.Foo, error)
	DeleteById(ctx context.Context, id string) error
}

type FooService interface {
	Create(ctx context.Context, foo *domain.Foo) (*domain.Foo, error)
	GetAll(ctx context.Context) ([]*domain.Foo, error)
	GetByID(ctx context.Context, id string) (*domain.Foo, error)
	GetByName(ctx context.Context, name string) (*domain.Foo, error)
	UpdateById(ctx context.Context, id string, update map[string]any) (*domain.Foo, error)
	DeleteById(ctx context.Context, id string) error
}
