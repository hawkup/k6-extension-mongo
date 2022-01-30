package mongo

import (
	"context"
	"time"

	"go.k6.io/k6/js/modules"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	modules.Register("k6/x/mongo", new(MONGO))
}

type MONGO struct{}

func (*MONGO) NewClient(connectionURI string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionURI))
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (*MONGO) DeleteOne(client *mongo.Client, dbName string, collName string, filter interface{}) (*mongo.DeleteResult, error) {
	coll := client.Database(dbName).Collection(collName)
	result, err := coll.DeleteOne(context.Background(), filter)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (*MONGO) FindOne(client *mongo.Client, dbName string, collName string, filter interface{}) *mongo.SingleResult {
	coll := client.Database(dbName).Collection(collName)
	result := coll.FindOne(context.Background(), filter)

	return result
}

func (*MONGO) InsertOne(client *mongo.Client, dbName string, collName string, doc interface{}) (*mongo.InsertOneResult, error) {
	coll := client.Database(dbName).Collection(collName)
	result, err := coll.InsertOne(context.Background(), doc)

	if err != nil {
		return nil, err
	}

	return result, nil
}
