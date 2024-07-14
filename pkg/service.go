package pkg

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service struct {
	db *mongo.Client
}

func NewService(client *mongo.Client) *Service {

	return &Service{
		db: client,
	}
}

func (s *Service) CreateUserdata(name, email, gender string) (Database, error) {
	coll := s.db.Database("sample_restaurants").Collection("restaurants")
	// Creates two sample documents describing restaurants
	user := Database{
		Name:   name,
		Email:  email,
		Gender: gender,
	}
	// Inserts sample documents into the collection
	_, err := coll.InsertOne(context.TODO(), user)
	if err != nil {
		log.Println("Error inserting user into MongoDB:", err)
		return Database{}, err
	}

	return user, nil
}
func (s *Service) GetAllUsers() ([]Database, error) {
	coll := s.db.Database("sample_restaurants").Collection("restaurants")
	// Creates a query filter to match documents in which the "cuisine"
	// is "Italian"
	//filter := bson.D{{"cuisine", "Italian"}}
	// Retrieves documents that match the query filter
	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())
	// Unpacks the cursor into a slice
	var users []Database
	if err = cursor.All(context.TODO(), &users); err != nil {
		return nil, err
	}
	return users, nil
}
