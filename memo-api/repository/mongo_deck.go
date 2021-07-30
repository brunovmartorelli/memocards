package repository

import (
	"context"
	"log"
	"time"

	"github.com/brunovmartorelli/memo-api/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DeckSchema struct {
	ID          primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Cards       []MongoCard        `json:"cards"`
	Size        int                `json:"size"`
}

type MongoDeck struct {
	Database   string
	Collection string
	Client     *mongo.Client
}

func NewDeck(c *mongo.Client) DeckRepository {
	return &MongoDeck{
		Client:     c,
		Database:   "memo",
		Collection: "deck",
	}
}

func (d *MongoDeck) Get(ID string) (*DeckSchema, error) {
	return &DeckSchema{}, nil
}

func (d *MongoDeck) GetByName(name string) (*DeckSchema, error) {
	collection := d.Client.Database(d.Database).Collection(d.Collection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res := collection.FindOne(ctx, bson.M{
		"name": name,
	})
	deck := &DeckSchema{}
	if err := res.Decode(deck); err != nil {
		return nil, err
	}
	return deck, nil
}

func (d *MongoDeck) Create(deck domain.Deck) error {
	collection := d.Client.Database(d.Database).Collection(d.Collection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, deck)
	if err != nil {
		return err
	}
	log.Printf("%v", res)
	return nil
}
func (d *MongoDeck) Update(name string, deck domain.Deck) (int64, error) {
	collection := d.Client.Database(d.Database).Collection(d.Collection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{
		"name": name,
	}

	update := bson.M{
		"$set": deck,
	}

	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err
	}
	return res.ModifiedCount, nil
}
func (d *MongoDeck) Delete(name string) (int64, error) {
	collection := d.Client.Database(d.Database).Collection(d.Collection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	count, err := collection.DeleteOne(ctx, bson.M{
		"name": name,
	})

	return count.DeletedCount, err
}
