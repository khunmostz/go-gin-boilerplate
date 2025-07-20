package tests

import (
	"context"
	"errors"
	"testing"

	"go-gin-boilerplate/internal/domain"
	"go-gin-boilerplate/internal/service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockFooRepo struct {
	mock.Mock
}

func (m *MockFooRepo) Create(ctx context.Context, foo *domain.Foo) (*domain.Foo, error) {
	args := m.Called(ctx, foo)
	return args.Get(0).(*domain.Foo), args.Error(1)
}
func (m *MockFooRepo) GetAll(ctx context.Context) ([]*domain.Foo, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*domain.Foo), args.Error(1)
}
func (m *MockFooRepo) GetByID(ctx context.Context, id string) (*domain.Foo, error) {
	args := m.Called(ctx, id)
	// Handle nil pointer case properly
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Foo), args.Error(1)
}
func (m *MockFooRepo) GetByName(ctx context.Context, name string) (*domain.Foo, error) {
	args := m.Called(ctx, name)
	// Handle nil pointer case properly
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Foo), args.Error(1)
}
func (m *MockFooRepo) UpdateById(ctx context.Context, id string, update map[string]any) (*domain.Foo, error) {
	args := m.Called(ctx, id, update)
	// Handle nil pointer case properly
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Foo), args.Error(1)
}
func (m *MockFooRepo) DeleteById(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestFooService_Create(t *testing.T) {
	repo := new(MockFooRepo)
	svc := service.NewFooService(repo)
	ctx := context.Background()
	foo := &domain.Foo{ID: "1", Name: "Test Foo"}

	repo.On("Create", ctx, foo).Return(foo, nil)
	result, err := svc.Create(ctx, foo)
	assert.NoError(t, err)
	assert.Equal(t, foo, result)

	fooInvalid := &domain.Foo{ID: "2", Name: ""}
	result, err = svc.Create(ctx, fooInvalid)
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestFooService_GetAll(t *testing.T) {
	repo := new(MockFooRepo)
	svc := service.NewFooService(repo)
	ctx := context.Background()
	foos := []*domain.Foo{{ID: "1", Name: "A"}, {ID: "2", Name: "B"}}

	repo.On("GetAll", ctx).Return(foos, nil)
	result, err := svc.GetAll(ctx)
	assert.NoError(t, err)
	assert.Equal(t, foos, result)
}

func TestFooService_GetByID(t *testing.T) {
	repo := new(MockFooRepo)
	svc := service.NewFooService(repo)
	ctx := context.Background()
	foo := &domain.Foo{ID: "1", Name: "A"}

	repo.On("GetByID", ctx, "1").Return(foo, nil)
	result, err := svc.GetByID(ctx, "1")
	assert.NoError(t, err)
	assert.Equal(t, foo, result)

	result, err = svc.GetByID(ctx, "")
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestFooService_GetByName(t *testing.T) {
	repo := new(MockFooRepo)
	svc := service.NewFooService(repo)
	ctx := context.Background()
	foo := &domain.Foo{ID: "1", Name: "A"}

	repo.On("GetByName", ctx, "A").Return(foo, nil)
	result, err := svc.GetByName(ctx, "A")
	assert.NoError(t, err)
	assert.Equal(t, foo, result)

	result, err = svc.GetByName(ctx, "")
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestFooService_UpdateById(t *testing.T) {
	repo := new(MockFooRepo)
	svc := service.NewFooService(repo)
	ctx := context.Background()
	foo := &domain.Foo{ID: "1", Name: "A"}
	update := map[string]any{"name": "B"}

	repo.On("GetByID", ctx, "1").Return(foo, nil)
	repo.On("UpdateById", ctx, "1", update).Return(&domain.Foo{ID: "1", Name: "B"}, nil)
	result, err := svc.UpdateById(ctx, "1", update)
	assert.NoError(t, err)
	assert.Equal(t, "B", result.Name)

	// Fixed: Use (*domain.Foo)(nil) instead of nil for proper type handling
	repo.On("GetByID", ctx, "2").Return((*domain.Foo)(nil), errors.New("not found"))
	result, err = svc.UpdateById(ctx, "2", update)
	assert.Error(t, err)
	assert.Nil(t, result)

	result, err = svc.UpdateById(ctx, "", update)
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestFooService_DeleteById(t *testing.T) {
	repo := new(MockFooRepo)
	svc := service.NewFooService(repo)
	ctx := context.Background()
	foo := &domain.Foo{ID: "1", Name: "A"}

	repo.On("GetByID", ctx, "1").Return(foo, nil)
	repo.On("DeleteById", ctx, "1").Return(nil)
	assert.NoError(t, svc.DeleteById(ctx, "1"))

	// Fixed: Use (*domain.Foo)(nil) instead of nil for proper type handling
	repo.On("GetByID", ctx, "2").Return((*domain.Foo)(nil), errors.New("not found"))
	assert.Error(t, svc.DeleteById(ctx, "2"))

	assert.Error(t, svc.DeleteById(ctx, ""))
}

// Example filter test for BaseRepository-like behavior
func TestFooService_Filter(t *testing.T) {
	repo := new(MockFooRepo)
	svc := service.NewFooService(repo)
	ctx := context.Background()
	foos := []*domain.Foo{{ID: "1", Name: "A"}, {ID: "2", Name: "B"}}

	repo.On("GetAll", ctx).Return(foos, nil)
	result, err := svc.GetAll(ctx)
	assert.NoError(t, err)
	var filtered []*domain.Foo
	for _, f := range result {
		if f.Name == "A" {
			filtered = append(filtered, f)
		}
	}
	assert.Len(t, filtered, 1)
	assert.Equal(t, "A", filtered[0].Name)
}
