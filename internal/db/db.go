package db

import "context"

type BaseRepository interface {
	Create(ctx context.Context, collection string, model any) error
	GetById(ctx context.Context, collection string, id string, result any) error
	GetAll(ctx context.Context, collection string, result any) error
	GetByField(ctx context.Context, collection string, field string, value any, result any) error
	UpdateById(ctx context.Context, collection string, id string, update any) error
	DeleteById(ctx context.Context, collection string, id string) error
}
