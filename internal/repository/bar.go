package repository

import (
	"context"
	"go-gin-boilerplate/internal/db"
	"go-gin-boilerplate/internal/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BarRepository struct {
	baseRepo   db.BaseRepository
	collection string
}

func NewBarRepository(baseRepo db.BaseRepository, collection string) *BarRepository {
	return &BarRepository{baseRepo: baseRepo, collection: collection}
}

func (br *BarRepository) CreateBar(ctx context.Context, bar *domain.Bar) (*domain.Bar, error) {
	bar.ID = primitive.NewObjectID().Hex()
	err := br.baseRepo.Create(ctx, br.collection, bar)
	if err != nil {
		return nil, err
	}
	return bar, nil
}

func (br *BarRepository) GetBars(ctx context.Context) ([]*domain.Bar, error) {
	var bars []*domain.Bar
	if err := br.baseRepo.GetAll(ctx, br.collection, &bars); err != nil {
		return nil, err
	}
	return bars, nil
}

func (br *BarRepository) GetBarByID(ctx context.Context, id string) (*domain.Bar, error) {
	var bar domain.Bar
	if err := br.baseRepo.GetById(ctx, br.collection, id, &bar); err != nil {
		return nil, err
	}
	return &bar, nil
}

func (br *BarRepository) UpdateBar(ctx context.Context, id string, bar *domain.Bar) (*domain.Bar, error) {
	err := br.baseRepo.UpdateById(ctx, br.collection, id, bar)
	if err != nil {
		return nil, err
	}
	// Set the ID for the updated bar
	bar.ID = id
	return bar, nil
}

func (br *BarRepository) DeleteBar(ctx context.Context, id string) error {
	return br.baseRepo.DeleteById(ctx, br.collection, id)
}
