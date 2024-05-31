package mongo

import (
	"context"

	"trainee3/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IMongoDB interface {
	GetDB() *mongo.Database
}

type mongoDB struct {
	db *mongo.Database
}

func New(cfg *config.Config) (IMongoDB, error) {
	clientOptions := options.Client().ApplyURI(cfg.MongoConfig.ApplyURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}
	if err := client.Ping(context.TODO(), nil); err != nil {
		return nil, err
	}

	return mongoDB{client.Database(cfg.MongoConfig.Database)}, nil
}

func (mdb mongoDB) GetDB() *mongo.Database {
	return mdb.db
}
