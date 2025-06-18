package db

import (
	"context"
	"go-gin-boilerplate/config"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongo(config *config.DatabaseConfig) *mongo.Client {
	safeDeref := func(s *string, defaultVal string) string {
		if s != nil {
			return *s
		}
		return defaultVal
	}
	dsn := safeDeref(config.URI, "mongodb://localhost:27017")
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
	return mg.client.Database(mg.dbName).Collection(collection).FindOne(ctx, bson.M{"_id": id}).Decode(result)
}

func (mg *mongoRepo) GetAll(ctx context.Context, collection string, result any) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	cur, err := mg.client.Database(mg.dbName).Collection(collection).Find(ctx, bson.M{})
	if err != nil {
		return err
	}
	return cur.All(ctx, result)
}

func (mg *mongoRepo) GetByField(ctx context.Context, collection string, field string, value any, result any) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	cur, err := mg.client.Database(mg.dbName).Collection(collection).Find(ctx, bson.M{field: value})
	if err != nil {
		return err
	}
	return cur.All(ctx, result)
}

func (mg *mongoRepo) UpdateById(ctx context.Context, collection string, id string, update any) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	_, err := mg.client.Database(mg.dbName).Collection(collection).UpdateOne(ctx, bson.M{"_id": id}, update)
	return err
}

func (mg *mongoRepo) DeleteById(ctx context.Context, collection string, id string) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	_, err := mg.client.Database(mg.dbName).Collection(collection).DeleteOne(ctx, bson.M{"_id": id})
	return err
}
