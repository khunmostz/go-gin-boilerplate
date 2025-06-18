package port

import (
	"context"
	"go-gin-boilerplate/internal/domain"
)

type FooRepository interface {
	CreateFoo(ctx context.Context, foo *domain.Foo) (*domain.Foo, error)
}

type FooService interface {
	CreateFoo(ctx context.Context, foo *domain.Foo) (*domain.Foo, error)
}
