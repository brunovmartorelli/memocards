package repository

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	Client *mongo.Client
}

func NewMongo() *Mongo {
	c, err := mongo.NewClient(options.Client().ApplyURI("<MONGODB_URI>"))
	if err != nil {
		log.Fatal(err)
	}
	return &Mongo{
		Client: c,
	}
}
