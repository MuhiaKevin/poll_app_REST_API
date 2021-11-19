package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Options struct {
	Description string `bson:"descsription"`
	Votes       int    `bson:"votes"`
}

type Poll struct {
	ID          primitive.ObjectID `bson:"_id"`
	PollID      string             `bson:"poll_id"`
	Total_votes int                `bson:"total_votes"`
	Question    string             `bson:"question"`
	Options     []Options          `bson:"options"`
}

var collection *mongo.Collection
var ctx = context.TODO()

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("polls").Collection("polls")
}

func CreatePoll(poll *Poll) error {
	_, err := collection.InsertOne(ctx, poll)
	return err
}

func GetAllPolls() ([]*Poll, error) {
	filter := bson.D{{}}
	return filterPolls(filter)
}

func filterPolls(filter interface{}) ([]*Poll, error) {
	var polls []*Poll

	cur, err := collection.Find(ctx, filter)

	if err != nil {
		return polls, err
	}

	for cur.Next(ctx) {
		var p Poll
		err := cur.Decode(&p)

		if err != nil {
			return polls, err
		}

		polls = append(polls, &p)
	}

	if err := cur.Err(); err != nil {
		return polls, err
	}

	cur.Close(ctx)

	if len(polls) == 0 {
		return polls, mongo.ErrNoDocuments
	}

	return polls, nil
}

func Find(id string) bson.M {
	var result bson.M

	err := collection.FindOne(ctx, bson.D{{Key: "poll_id", Value: id}}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println(err)
		}
		log.Fatal(err)
	}
	for _, value := range result {
		fmt.Printf("%v\n", value)
	}

	return result
}

// func UpdateVote(id string, optionIndex int) {

// }
