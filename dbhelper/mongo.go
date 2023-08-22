package dbhelper

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConn(ctx context.Context, applyUrl, database string) *mongo.Database {
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(applyUrl))
	if err != nil {
		panic(err)
	}

	return mongoClient.Database(database)
}
