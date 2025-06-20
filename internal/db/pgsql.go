package db

import (
	"context"
	"errors"
	"fmt"
	"go-gin-boilerplate/config"
	"go-gin-boilerplate/internal/domain"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitPgsql(config *config.DatabaseConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s timezone=%s",
		config.Postgres.Host,
		config.Postgres.Port,
		config.Postgres.Username,
		config.Postgres.Password,
		config.Postgres.DBName,
		config.Postgres.SSLMode,
		config.Postgres.Timezone,
	)

	gormConfig := &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Info),
	}
	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}

	log.Println("Successfully connected to PostgreSQL!")

	db.AutoMigrate(&domain.Foo{}, &domain.Bar{})

	return db
}

type pgsqlRepository struct {
	db *gorm.DB
}

func NewPgsqlRepository(db *gorm.DB) BaseRepository {
	return &pgsqlRepository{db: db}
}

func (pg *pgsqlRepository) Create(ctx context.Context, _ string, model any) error {
	return pg.db.WithContext(ctx).Create(model).Error
}

func (pg *pgsqlRepository) GetById(ctx context.Context, _ string, id string, result any) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	err := pg.db.WithContext(ctx).First(result, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("entity not found")
		}
		return err
	}
	return nil
}

func (pg *pgsqlRepository) GetAll(ctx context.Context, _ string, result any) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	return pg.db.WithContext(ctx).Find(result).Error
}

func (pg *pgsqlRepository) GetByField(ctx context.Context, _ string, field string, value any, result any) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	condition := map[string]any{field: value}
	res := pg.db.WithContext(ctx).Where(condition).First(result)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return errors.New("entity not found")
		}
		return res.Error
	}
	return nil
}

func (pg *pgsqlRepository) UpdateById(ctx context.Context, collection string, id string, update any) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	model, err := pg.getModelByCollection(collection)
	if err != nil {
		return err
	}

	result := pg.db.WithContext(ctx).Model(model).Where("id = ?", id).Updates(update)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("entity not found")
		}
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("entity not found")
	}

	return nil
}

func (pg *pgsqlRepository) DeleteById(ctx context.Context, collection string, id string) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	model, err := pg.getModelByCollection(collection)
	if err != nil {
		return err
	}

	result := pg.db.WithContext(ctx).Delete(model, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("entity not found")
	}

	return nil
}

// getModelByCollection returns the appropriate model based on collection name
func (pg *pgsqlRepository) getModelByCollection(collection string) (any, error) {
	switch collection {
	case "bar":
		return &domain.Bar{}, nil
	case "foo":
		return &domain.Foo{}, nil
	default:
		return nil, fmt.Errorf("unknown collection: %s", collection)
	}
}
