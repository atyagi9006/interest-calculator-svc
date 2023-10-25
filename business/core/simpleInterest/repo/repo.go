package repo

import (
	"context"

	"github.com/atyagi9006/interest-calculator-svc/business/core/entities"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Create(ctx context.Context, simpleinterest *entities.SimpleInterest) (*entities.SimpleInterest, error)
	Read(ctx context.Context,id string) (*entities.SimpleInterest, error)
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
	return nil, nil
}
func (repo *repository) Read(ctx context.Context,id string) (*entities.SimpleInterest, error) {
	return nil, nil
}
func (repo *repository) Delete(ctx context.Context, ID string) error {
	return nil
}
