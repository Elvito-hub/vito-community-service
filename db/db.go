package db

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoClient *mongo.Client
)

func GetMongoClient() (*mongo.Client, error) {

	dbURL := os.Getenv("dbUrl")

	clientOptions := options.Client().ApplyURI(dbURL)

	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		return nil, err
	}

	// check if the connection was successful
	err = client.Ping(context.Background(), nil)

	if err != nil {
		return nil, err
	}

	MongoClient = client

	return client, nil
}
