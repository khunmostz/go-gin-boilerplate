package service

import (
	"context"
	"errors"
	"go-gin-boilerplate/internal/domain"
	"go-gin-boilerplate/internal/port"
	"strings"
)

type BarService struct {
	barRepo port.BarRepository
}

func NewBarService(barRepo port.BarRepository) port.BarService {
	return &BarService{barRepo: barRepo}
}

func (bs *BarService) Create(ctx context.Context, bar *domain.Bar) (*domain.Bar, error) {
	// Validate required fields
	if strings.TrimSpace(bar.Name) == "" {
		return nil, errors.New("bar name is required")
	}

	// Set default status if not provided
	if strings.TrimSpace(bar.Status) == "" {
		bar.Status = "active"
	}

	return bs.barRepo.Create(ctx, bar)
}

func (bs *BarService) GetAll(ctx context.Context) ([]*domain.Bar, error) {
	return bs.barRepo.GetAll(ctx)
}

func (bs *BarService) GetByID(ctx context.Context, id string) (*domain.Bar, error) {
	if strings.TrimSpace(id) == "" {
		return nil, errors.New("bar ID is required")
	}
	return bs.barRepo.GetByID(ctx, id)
}

func (bs *BarService) GetByName(ctx context.Context, name string) (*domain.Bar, error) {
	if strings.TrimSpace(name) == "" {
		return nil, errors.New("bar name is required")
	}
	return bs.barRepo.GetByName(ctx, name)
}

func (bs *BarService) UpdateById(ctx context.Context, id string, update map[string]any) (*domain.Bar, error) {
	if strings.TrimSpace(id) == "" {
		return nil, errors.New("bar ID is required")
	}

	// Check if bar exists
	_, err := bs.barRepo.GetByID(ctx, id)
	if err != nil {
		return nil, errors.New("bar not found")
	}

	return bs.barRepo.UpdateById(ctx, id, update)
}

func (bs *BarService) DeleteById(ctx context.Context, id string) error {
	if strings.TrimSpace(id) == "" {
		return errors.New("bar ID is required")
	}

	// Check if bar exists
	_, err := bs.barRepo.GetByID(ctx, id)
	if err != nil {
		return errors.New("bar not found")
	}

	return bs.barRepo.DeleteById(ctx, id)
}
