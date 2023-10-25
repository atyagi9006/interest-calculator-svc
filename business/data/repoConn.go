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

func (rc *RepoConn) NewRepoConn() (*mongo.Database, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		rc.MongoURL).SetServerSelectionTimeout(time.Duration(rc.SelectionTimeOut)*time.
		Second))
	if err != nil {
		cancel()
		return nil, nil, err
	}
	db := client.Database(rc.MongoDBName)
	return db, cancel, nil
}
