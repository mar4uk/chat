package model

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBClient struct {
	*mongo.Client
}

type DB struct {
	*mongo.Database
}

func NewClient(uri string) (*DBClient, error) {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://root:password@localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, err
	}

	// Check the connection
	if err = client.Ping(context.TODO(), nil); err != nil {
		return nil, err
	}

	return &DBClient{client}, nil
}
