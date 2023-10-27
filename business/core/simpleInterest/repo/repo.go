package repo

import (
	"context"

	"github.com/atyagi9006/interest-calculator-svc/business/core/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//Repository is a interface to interact and do operation on mongodb
type Repository interface {
	Create(ctx context.Context, simpleinterest *entities.SimpleInterest) (*entities.SimpleInterest, error)
	Read(ctx context.Context, id string) (*entities.SimpleInterest, error)
	Delete(ctx context.Context, ID string) error
}

//repository handle and implements the operation on mongodb  
type repository struct {
	Collection *mongo.Collection
}

// NewRepo is the single instance repo that is being created.
func NewRepo(collection *mongo.Collection) Repository {
	return &repository{
		Collection: collection,
	}
}

//Create it creates the simpleinterest data on mongodb
func (repo *repository) Create(ctx context.Context, simpleinterest *entities.SimpleInterest) (*entities.SimpleInterest, error) {

	_, err := repo.Collection.InsertOne(ctx, simpleinterest)
	if err != nil {
		return nil, err
	}
	return simpleinterest, nil
}

//Read it reads the simpleinterest data on mongodb
func (repo *repository) Read(ctx context.Context, id string) (*entities.SimpleInterest, error) {

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	res := &entities.SimpleInterest{}

	filter := bson.D{{Key: "_id", Value: _id}}
	if err := repo.Collection.FindOne(ctx, filter).Decode(&res); err != nil {
		return nil, err
	}
	return res, nil
}

//Delete it deletes the simpleinterest data on mongodb
func (repo *repository) Delete(ctx context.Context, ID string) error {
	simpleinterestID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = repo.Collection.DeleteOne(ctx, bson.M{"_id": simpleinterestID})
	if err != nil {
		return err
	}
	return nil
}
