package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/brunovmartorelli/memo-api/domain/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CardSchema struct {
	Front string `json:"front"`
	Back  string `json:"back"`
	Score int    `json:"score"`
}

type MongoCard struct {
	Database   string
	Collection string
	Client     *mongo.Client
}

func NewCard(c *mongo.Client) CardRepository {
	return &MongoCard{
		Client:     c,
		Database:   "memo",
		Collection: "deck",
	}
}

func (c *MongoCard) GetByFront(front, deckName string) (*CardSchema, error) {
	collection := c.Client.Database(c.Database).Collection(c.Collection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{
		"name": deckName,
		"cards": bson.M{
			"$elemMatch": bson.M{
				"front": front,
			},
		},
	}
	res := collection.FindOne(ctx, filter)
	card := CardSchema{}
	if err := res.Decode(&card); err != nil {
		return nil, err
	}
	return &card, nil
}

func (c *MongoCard) List(deckName string) (*[]CardSchema, error) {
	collection := c.Client.Database(c.Database).Collection(c.Collection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{
		"name": deckName,
	}

	options := options.FindOne()
	options.SetProjection(bson.M{
		"cards": 1,
		"_id":   0,
	})

	res := collection.FindOne(ctx, filter, options)

	deck := &DeckSchema{}
	if err := res.Decode(deck); err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(deck)
	cards := &deck.Cards
	return cards, nil
}
func (c *MongoCard) Create(deckName string, card entities.Card) error {
	collection := c.Client.Database(c.Database).Collection(c.Collection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{
		"name": deckName,
	}

	update := bson.M{
		"$push": bson.M{
			"cards": card,
		},
	}

	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if res.ModifiedCount < 1 {
		return fmt.Errorf("failed to create card")
	}

	return nil
}
func (c *MongoCard) Update(front, deckName string, card entities.Card) (int64, error) {
	collection := c.Client.Database(c.Database).Collection(c.Collection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{
		"name": deckName,
		"cards": bson.M{
			"$elemMatch": bson.M{
				"front": front,
			},
		},
	}

	update := bson.M{
		"$set": bson.M{"cards.$": card},
	}

	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err
	}
	return res.ModifiedCount, err
}
func (c *MongoCard) Delete(front, deckName string) (int64, error) {
	collection := c.Client.Database(c.Database).Collection(c.Collection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{
		"name": deckName,
	}
	update := bson.M{
		"$pull": bson.M{
			"cards": bson.M{
				"front": front,
			},
		},
	}

	count, err := collection.UpdateOne(ctx, filter, update)
	return count.ModifiedCount, err
}
