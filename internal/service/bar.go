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

func (bs *BarService) CreateBar(ctx context.Context, bar *domain.Bar) (*domain.Bar, error) {
	// Validate required fields
	if strings.TrimSpace(bar.Name) == "" {
		return nil, errors.New("bar name is required")
	}

	// Set default status if not provided
	if strings.TrimSpace(bar.Status) == "" {
		bar.Status = "active"
	}

	return bs.barRepo.CreateBar(ctx, bar)
}

func (bs *BarService) GetBars(ctx context.Context) ([]*domain.Bar, error) {
	return bs.barRepo.GetBars(ctx)
}

func (bs *BarService) GetBarByID(ctx context.Context, id string) (*domain.Bar, error) {
	if strings.TrimSpace(id) == "" {
		return nil, errors.New("bar ID is required")
	}
	return bs.barRepo.GetBarByID(ctx, id)
}

func (bs *BarService) UpdateBar(ctx context.Context, id string, bar *domain.Bar) (*domain.Bar, error) {
	if strings.TrimSpace(id) == "" {
		return nil, errors.New("bar ID is required")
	}

	// Validate required fields
	if strings.TrimSpace(bar.Name) == "" {
		return nil, errors.New("bar name is required")
	}

	// Check if bar exists
	_, err := bs.barRepo.GetBarByID(ctx, id)
	if err != nil {
		return nil, errors.New("bar not found")
	}

	return bs.barRepo.UpdateBar(ctx, id, bar)
}

func (bs *BarService) DeleteBar(ctx context.Context, id string) error {
	if strings.TrimSpace(id) == "" {
		return errors.New("bar ID is required")
	}

	// Check if bar exists
	_, err := bs.barRepo.GetBarByID(ctx, id)
	if err != nil {
		return errors.New("bar not found")
	}

	return bs.barRepo.DeleteBar(ctx, id)
}
