package database

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store struct {
	Collection *mongo.Collection
}

func NewClient(uri string) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("NewClient, NewClient: %w", err)
	}

	return client, nil
}

func (s *Store) InitStore(client *mongo.Client) error {
	return nil
}
