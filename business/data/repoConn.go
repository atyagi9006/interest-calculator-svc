package data

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//RepoConn provide repo connection to mongo
type RepoConn struct {
	MongoURL         string
	SelectionTimeOut int
	MongoDBName      string
}

//NewRepoConn initialte and returns a mongodb client
func (rc *RepoConn) NewRepoConn(ctx context.Context) (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		rc.MongoURL).SetServerSelectionTimeout(time.Duration(rc.SelectionTimeOut)*time.
		Second))
	if err != nil {
		
		return nil, err
	}
	return client, nil
}
