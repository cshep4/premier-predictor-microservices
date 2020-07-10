package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func New(ctx context.Context) (*mongo.Client, error) {
	opts := options.Client().
		SetConnectTimeout(2 * time.Second).
		SetServerSelectionTimeout(2 * time.Second).
		ApplyURI(uri())

	err := opts.Validate()
	if err != nil {
		return nil, fmt.Errorf("opts_validate: %v", err)
	}

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("connect: %v", err)
	}

	return client, nil
}

func uri() string {
	username := os.Getenv("MONGO_USERNAME")
	password := os.Getenv("MONGO_PASSWORD")
	port := os.Getenv("MONGO_PORT")

	uri := fmt.Sprintf("%s://", os.Getenv("MONGO_SCHEME"))
	if username != "" && password != "" {
		uri = fmt.Sprintf("%s%s:%s@", uri, username, password)
	}
	uri = uri + os.Getenv("MONGO_HOST")
	if port != "" {
		uri = fmt.Sprintf("%s:%s", uri, port)
	}

	return uri
}
