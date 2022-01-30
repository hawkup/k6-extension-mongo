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
