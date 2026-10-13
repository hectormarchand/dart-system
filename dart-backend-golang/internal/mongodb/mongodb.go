package mongodb

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const databaseName = "dart"

// Connect dials MongoDB using the "mongodbUrl" env var (falls back to the local dev
// instance from docker/local/docker-compose.yml) and returns the connected client
// along with the "dart" database handle.
func Connect(ctx context.Context) (*mongo.Client, *mongo.Database, error) {
	uri := os.Getenv("mongodbUrl")
	if uri == "" {
		uri = "mongodb://dart:d4rt@127.0.0.1:27017"
	}

	clientOptions := options.Client().
		ApplyURI(uri)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, nil, err
	}

	return client, client.Database(databaseName), nil
}
