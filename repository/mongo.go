package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type IMongo interface {
	FindOne(db *mongo.Database, collectionStr string, filter interface{}, model interface{}) error
	Insert(db *mongo.Database, collectionStr string, values interface{}) error
}

type Mongo struct {
	db *mongo.Database
}

func NewMongo() IMongo {
	return &Mongo{}

}

func (r *Mongo) FindOne(db *mongo.Database, collectionStr string, filter interface{}, model interface{}) error {
	collection := db.Collection(collectionStr)
	if err := collection.FindOne(context.Background(), filter).Decode(model); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		return err
	}
	return nil
}

func (r *Mongo) Insert(db *mongo.Database, collectionStr string, values interface{}) error {
	collection := db.Collection(collectionStr)
	if _, err := collection.InsertOne(context.Background(), values); err != nil {
		return err
	}
	return nil
}
