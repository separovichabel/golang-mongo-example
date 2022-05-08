package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BaseCRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func (repository *BaseCRepository) Create() (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	res, err := repository.collection.InsertOne(ctx, bson.D{{"name", "pi"}, {"value", 3.14159}})

	if err != nil {
		return nil, err
	}

	return res.InsertedID, nil
}

func (repository *BaseCRepository) FindAll() (*[]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	res, err := repository.collection.Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}
	resp := new([]interface{})
	res.All(ctx, resp)
	fmt.Println("CONTENT: ", resp)
	return resp, nil
}

func NewBaseCRepository(client *mongo.Client, collection *mongo.Collection) *BaseCRepository {
	return &BaseCRepository{client: client, collection: collection}
}

func ConnectDatabase(config *Config) (*mongo.Client, *mongo.Collection, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.DatabasePath))

	if err != nil {
		return nil, nil, err
	}

	collection := client.Database(config.DatabaseName).Collection(config.CollectionName)

	return client, collection, nil
}
