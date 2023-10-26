package data

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RepoConn struct {
	MongoURL         string
	SelectionTimeOut int
	MongoDBName      string
}

func (rc *RepoConn) NewRepoConn(ctx context.Context) (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		rc.MongoURL).SetServerSelectionTimeout(time.Duration(rc.SelectionTimeOut)*time.
		Second))
	if err != nil {
		
		return nil, err
	}
	return client, nil
}
