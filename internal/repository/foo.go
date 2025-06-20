package repository

import (
	"context"
	"go-gin-boilerplate/internal/db"
	"go-gin-boilerplate/internal/domain"
	"go-gin-boilerplate/internal/port"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FooRepository struct {
	baseRepo   db.BaseRepository
	collection string
}

func NewFooRepository(baseRepo db.BaseRepository, collection string) port.FooRepository {
	return &FooRepository{baseRepo: baseRepo, collection: collection}
}

func (fr *FooRepository) Create(ctx context.Context, foo *domain.Foo) (*domain.Foo, error) {
	foo.ID = primitive.NewObjectID().Hex()
	err := fr.baseRepo.Create(ctx, fr.collection, foo)
	if err != nil {
		return nil, err
	}
	return foo, nil
}

func (fr *FooRepository) GetAll(ctx context.Context) ([]*domain.Foo, error) {
	var foos []*domain.Foo
	if err := fr.baseRepo.GetAll(ctx, fr.collection, &foos); err != nil {
		return nil, err
	}
	return foos, nil
}

func (fr *FooRepository) GetByID(ctx context.Context, id string) (*domain.Foo, error) {
	var foo domain.Foo
	if err := fr.baseRepo.GetById(ctx, fr.collection, id, &foo); err != nil {
		return nil, err
	}
	return &foo, nil
}

func (fr *FooRepository) GetByName(ctx context.Context, name string) (*domain.Foo, error) {
	var foo domain.Foo
	if err := fr.baseRepo.GetByField(ctx, fr.collection, "name", name, &foo); err != nil {
		return nil, err
	}
	return &foo, nil
}

func (fr *FooRepository) UpdateById(ctx context.Context, id string, update map[string]any) (*domain.Foo, error) {
	err := fr.baseRepo.UpdateById(ctx, fr.collection, id, update)
	if err != nil {
		return nil, err
	}

	// Get the updated foo from database to return the complete object
	var updatedFoo *domain.Foo
	err = fr.baseRepo.GetById(ctx, fr.collection, id, &updatedFoo)
	if err != nil {
		return nil, err
	}

	return updatedFoo, nil
}

func (fr *FooRepository) DeleteById(ctx context.Context, id string) error {
	return fr.baseRepo.DeleteById(ctx, fr.collection, id)
}
