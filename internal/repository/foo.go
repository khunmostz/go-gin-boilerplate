package repository

import (
	"context"
	"go-gin-boilerplate/internal/db"
	"go-gin-boilerplate/internal/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FooRepository struct {
	baseRepo   db.BaseRepository
	collection string
}

func NewFooRepository(baseRepo db.BaseRepository, collection string) *FooRepository {
	return &FooRepository{baseRepo: baseRepo, collection: collection}
}

func (fr *FooRepository) CreateFoo(ctx context.Context, foo *domain.Foo) (*domain.Foo, error) {
	foo.ID = primitive.NewObjectID().Hex()
	err := fr.baseRepo.Create(ctx, fr.collection, foo)
	if err != nil {
		return nil, err
	}
	return foo, nil
}
