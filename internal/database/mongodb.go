package database

import (
	"context"
	"fmt"

	"github.com/madxiii/mongocrud/internal/models"
	"go.mongodb.org/mongo-driver/bson"
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

func (s *Store) InitCollection(client *mongo.Client) error {
	s.Collection = client.Database("post").Collection("users")
	return nil
}

func (s *Store) Find(c context.Context) ([]models.User, error) {
	var users []bson.M

	filter := bson.M{}
	cursor, err := s.Collection.Find(c, filter)
	if err != nil {
		return nil, fmt.Errorf("database Find, Find: %w", err)
	}

	cursor.All(c, &users)

	fmt.Println(&users)
	return nil, nil
}

func (s *Store) Insert(c context.Context, user models.User) error {
	res, err := s.Collection.Indexes().CreateOne(c, mongo.IndexModel{Keys: bson.D{{Key: "email", Value: 1}}, Options: options.Index().SetUnique(true)})
	if err != nil {
		return fmt.Errorf("database Insert, CreateOne: %w", err)
	}
	fmt.Println("Res:", res)
	_, err = s.Collection.InsertOne(c, user)
	if err != nil {
		return fmt.Errorf("database Insert, InsertOne: %w", err)
	}
	// id, ok := res.InsertedID.(primitive.ObjectID)
	// if ok {
	// 	fmt.Println("true", id)
	// 	return nil
	// }
	// fmt.Println("false", id)
	return nil
}

func (s *Store) Update(c context.Context, user models.User) error {
	return nil
}

func (s *Store) Delete(c context.Context, user models.User) error {
	return nil
}
