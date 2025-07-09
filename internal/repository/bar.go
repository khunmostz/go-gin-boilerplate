package repository

import (
	"context"
	"go-gin-boilerplate/internal/db"
	"go-gin-boilerplate/internal/domain"
	"go-gin-boilerplate/internal/port"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BarRepository struct {
	baseRepo   db.BaseRepository
	collection string
}

func NewBarRepository(baseRepo db.BaseRepository, collection string) port.BarRepository {
	return &BarRepository{baseRepo: baseRepo, collection: collection}
}

func (br *BarRepository) Create(ctx context.Context, bar *domain.Bar) (*domain.Bar, error) {
	bar.ID = primitive.NewObjectID().Hex()
	err := br.baseRepo.Create(ctx, br.collection, bar)
	if err != nil {
		return nil, err
	}
	return bar, nil
}

func (br *BarRepository) GetAll(ctx context.Context) ([]*domain.Bar, error) {
	var bars []*domain.Bar
	if err := br.baseRepo.GetAll(ctx, br.collection, &bars, nil); err != nil {
		return nil, err
	}
	return bars, nil
}

func (br *BarRepository) GetByID(ctx context.Context, id string) (*domain.Bar, error) {
	var bar domain.Bar
	if err := br.baseRepo.GetById(ctx, br.collection, id, &bar); err != nil {
		return nil, err
	}
	return &bar, nil
}

func (br *BarRepository) GetByName(ctx context.Context, name string) (*domain.Bar, error) {
	var bar domain.Bar
	if err := br.baseRepo.GetByField(ctx, br.collection, "name", name, &bar); err != nil {
		return nil, err
	}
	return &bar, nil
}

func (br *BarRepository) UpdateById(ctx context.Context, id string, update map[string]any) (*domain.Bar, error) {
	err := br.baseRepo.UpdateById(ctx, br.collection, id, update)
	if err != nil {
		return nil, err
	}

	// Get the updated bar from database to return the complete object
	var updatedBar *domain.Bar
	err = br.baseRepo.GetById(ctx, br.collection, id, &updatedBar)
	if err != nil {
		return nil, err
	}

	return updatedBar, nil
}

func (br *BarRepository) DeleteById(ctx context.Context, id string) error {
	return br.baseRepo.DeleteById(ctx, br.collection, id)
}
