package database

import (
	"context"
	"fmt"

	"github.com/madxiii/mongocrud/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (s *Store) Find(c context.Context) ([]primitive.M, error) {
	var users []bson.M

	filter := bson.M{}
	cursor, err := s.Collection.Find(c, filter)
	if err != nil {
		return nil, fmt.Errorf("database Find, Find: %w", err)
	}

	cursor.All(c, &users)

	return users, nil
}

func (s *Store) Insert(c context.Context, user models.User) error {
	_, err := s.Collection.Indexes().CreateMany(c, []mongo.IndexModel{
		{Keys: bson.D{{Key: "nickname", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "email", Value: 1}}, Options: options.Index().SetUnique(true)},
	})
	if err != nil {
		return fmt.Errorf("database Insert, CreateOne: %w", err)
	}
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

func (s *Store) Update(c context.Context, id string, newData models.User) error {
	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("datatbase Update, ObjectIDFromHex: %w", err)
	}
	filter := bson.M{"_id": bson.M{"$eq": docID}}
	update := bson.M{
		"$set": bson.M{
			"nickname": newData.Nickname,
		},
		// {"email": newData.Email},
		// {"name": newData.Name},
		// {"age": newData.Age},
	}

	res, err := s.Collection.UpdateByID(c, filter, update)
	if err != nil {
		return fmt.Errorf("database Update, UpdateOne: %w", err)
	}
	fmt.Println(res)
	return nil
}

func (s *Store) Delete(c context.Context, id string) error {
	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("database Delete, ObjectIDFromHex: %w", err)
	}

	filter := bson.M{"_id": bson.M{"$eq": docID}}

	_, err = s.Collection.DeleteOne(c, filter)
	if err != nil {
		return fmt.Errorf("database Delete, DeleteOne: %w", err)
	}
	return nil
}
