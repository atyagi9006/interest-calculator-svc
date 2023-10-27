package repo

import (
	"context"
	"time"

	"github.com/atyagi9006/interest-calculator-svc/business/core/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Create(ctx context.Context, simpleinterest *entities.SimpleInterest) (*entities.SimpleInterest, error)
	Read(ctx context.Context, id string) (*entities.SimpleInterest, error)
	Delete(ctx context.Context, ID string) error
}
type repository struct {
	Collection *mongo.Collection
}

// NewRepo is the single instance repo that is being created.
func NewRepo(collection *mongo.Collection) Repository {
	return &repository{
		Collection: collection,
	}
}

func (repo *repository) Create(ctx context.Context, simpleinterest *entities.SimpleInterest) (*entities.SimpleInterest, error) {
	simpleinterest.ID = primitive.NewObjectID()
	simpleinterest.CreatedAt = time.Now()
	simpleinterest.UpdatedAt = time.Now()
	_, err := repo.Collection.InsertOne(ctx, simpleinterest)
	if err != nil {
		return nil, err
	}
	return simpleinterest, nil
}
func (repo *repository) Read(ctx context.Context, id string) (*entities.SimpleInterest, error) {

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil,err
	}
	res:=&entities.SimpleInterest{}

	filter := bson.D{{Key: "_id", Value: _id}}
	if err :=repo.Collection.FindOne(ctx, filter).Decode(&res); err != nil {
		return nil,err
	}
	return res, nil
}
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
