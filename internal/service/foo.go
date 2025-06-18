package service

import (
	"context"
	"go-gin-boilerplate/internal/domain"
	"go-gin-boilerplate/internal/port"
)

type FooService struct {
	fooRepo port.FooRepository
}

func NewFooService(fooRepo port.FooRepository) port.FooService {
	return &FooService{fooRepo: fooRepo}
}

func (fs *FooService) CreateFoo(ctx context.Context, foo *domain.Foo) (*domain.Foo, error) {
	return fs.fooRepo.CreateFoo(ctx, foo)
}
