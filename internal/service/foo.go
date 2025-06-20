package service

import (
	"context"
	"errors"
	"go-gin-boilerplate/internal/domain"
	"go-gin-boilerplate/internal/port"
	"strings"
)

type FooService struct {
	fooRepo port.FooRepository
}

func NewFooService(fooRepo port.FooRepository) port.FooService {
	return &FooService{fooRepo: fooRepo}
}

func (fs *FooService) Create(ctx context.Context, foo *domain.Foo) (*domain.Foo, error) {
	// Validate required fields
	if strings.TrimSpace(foo.Name) == "" {
		return nil, errors.New("foo name is required")
	}

	return fs.fooRepo.Create(ctx, foo)
}

func (fs *FooService) GetAll(ctx context.Context) ([]*domain.Foo, error) {
	return fs.fooRepo.GetAll(ctx)
}

func (fs *FooService) GetByID(ctx context.Context, id string) (*domain.Foo, error) {
	if strings.TrimSpace(id) == "" {
		return nil, errors.New("foo ID is required")
	}
	return fs.fooRepo.GetByID(ctx, id)
}

func (fs *FooService) GetByName(ctx context.Context, name string) (*domain.Foo, error) {
	if strings.TrimSpace(name) == "" {
		return nil, errors.New("foo name is required")
	}
	return fs.fooRepo.GetByName(ctx, name)
}

func (fs *FooService) UpdateById(ctx context.Context, id string, update map[string]any) (*domain.Foo, error) {
	if strings.TrimSpace(id) == "" {
		return nil, errors.New("foo ID is required")
	}

	// Check if foo exists
	_, err := fs.fooRepo.GetByID(ctx, id)
	if err != nil {
		return nil, errors.New("foo not found")
	}

	return fs.fooRepo.UpdateById(ctx, id, update)
}

func (fs *FooService) DeleteById(ctx context.Context, id string) error {
	if strings.TrimSpace(id) == "" {
		return errors.New("foo ID is required")
	}

	// Check if foo exists
	_, err := fs.fooRepo.GetByID(ctx, id)
	if err != nil {
		return errors.New("foo not found")
	}

	return fs.fooRepo.DeleteById(ctx, id)
}
