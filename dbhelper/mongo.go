package dbhelper

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoClient(ctx context.Context, applyUrl string) *mongo.Client {
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(applyUrl))
	if err != nil {
		panic(err)
	}

	return mongoClient
}
