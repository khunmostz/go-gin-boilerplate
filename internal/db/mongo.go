package db

import (
	"context"
	"errors"
	"go-gin-boilerplate/config"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongo(config *config.DatabaseConfig) *mongo.Client {
	dsn := config.MongoDB.URI
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(dsn))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	log.Println("Successfully connected to MongoDB!")
	return client
}

type mongoRepo struct {
	client *mongo.Client
	dbName string
}

func NewMongoRepository(client *mongo.Client, dbName string) BaseRepository {
	return &mongoRepo{client: client, dbName: dbName}
}

func (mg *mongoRepo) Create(ctx context.Context, collection string, model any) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	_, err := mg.client.Database(mg.dbName).Collection(collection).InsertOne(ctx, model)
	return err
}

func (mg *mongoRepo) GetById(ctx context.Context, collection string, id string, result any) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	err := mg.client.Database(mg.dbName).Collection(collection).FindOne(ctx, bson.M{"_id": id}).Decode(result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return errors.New("entity not found")
		}
		return err
	}
	return nil
}

func (mg *mongoRepo) GetAll(ctx context.Context, collection string, result any, filter any) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	mongoFilter := bson.M{}
	if filter != nil {
		// Use bson.Marshal and Unmarshal for automatic conversion
		filterBytes, err := bson.Marshal(filter)
		if err != nil {
			return err
		}
		err = bson.Unmarshal(filterBytes, &mongoFilter)
		if err != nil {
			return err
		}

		// Remove empty/zero values to avoid matching empty strings
		cleanFilter := bson.M{}
		for key, value := range mongoFilter {
			if value != nil && value != "" {
				cleanFilter[key] = value
			}
		}
		mongoFilter = cleanFilter
	}

	cur, err := mg.client.Database(mg.dbName).Collection(collection).Find(ctx, mongoFilter)
	if err != nil {
		return err
	}
	return cur.All(ctx, result)
}

func (mg *mongoRepo) GetByField(ctx context.Context, collection string, field string, value any, result any) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	err := mg.client.Database(mg.dbName).Collection(collection).FindOne(ctx, bson.M{field: value}).Decode(result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return errors.New("entity not found")
		}
		return err
	}
	return nil
}

func (mg *mongoRepo) UpdateById(ctx context.Context, collection string, id string, update any) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	result, err := mg.client.Database(mg.dbName).Collection(collection).UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": update})
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("entity not found")
	}
	return nil
}

func (mg *mongoRepo) DeleteById(ctx context.Context, collection string, id string) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	result, err := mg.client.Database(mg.dbName).Collection(collection).DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("entity not found")
	}
	return nil
}
