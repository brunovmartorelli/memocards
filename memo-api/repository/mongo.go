package repository

import (
	"context"
	"log"

	"github.com/brunovmartorelli/memo-api/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	Client *mongo.Client
	cfg    *config.Mongo
}

func NewMongo(cfg *config.Mongo) *Mongo {
	c, err := mongo.NewClient(options.Client().ApplyURI(cfg.URI))
	if err != nil {
		log.Fatal(err)
	}
	return &Mongo{
		Client: c,
		cfg:    cfg,
	}
}

func (m *Mongo) Disconnect(ctx context.Context) {
	if err := m.Client.Disconnect(ctx); err != nil {
		panic(err)
	}
}
