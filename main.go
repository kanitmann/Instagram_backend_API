package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type profile struct {
	ID     string `json:"ID" bson:"ID"`
	Name   string `json:"name" bson:"name"`
	Age    int    `json:"age" bson:"age"`
	Postid int    `json:"postid" bson:"postid"`
}

type total_users struct {
	Start []int
	End   []int
}

//To add values to map for keeping track of users
func (data *total_users) AppendValues(s int, e int) {
	data.Start = append(data.Start, s)
	data.End = append(data.End, s)
}

func close(client *mongo.Client, ctx context.Context,
	cancel context.CancelFunc) {

	defer cancel()

	defer func() {

		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func connect(uri string) (*mongo.Client, context.Context,
	context.CancelFunc, error) {

	ctx, cancel := context.WithTimeout(context.Background(),
		30*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

func ping(client *mongo.Client, ctx context.Context) error {

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	fmt.Println("connected successfully")
	return nil
}

func insertOne(client *mongo.Client, ctx context.Context, dataBase, col string, doc interface{}) (*mongo.InsertOneResult, error) {

	collection := client.Database(dataBase).Collection(col)

	result, err := collection.InsertOne(ctx, doc)
	return result, err
}

func insertMany(client *mongo.Client, ctx context.Context, dataBase, col string, docs []interface{}) (*mongo.InsertManyResult, error) {

	collection := client.Database(dataBase).Collection(col)
	result, err := collection.InsertMany(ctx, docs)
	return result, err
}

func main() {
	client, ctx, cancel, err := connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	defer close(client, ctx, cancel)

	quickstartDatabase := client.Database("Quickstart")
	TaskCollection := quickstartDatabase.Collection("Task")
	episodesCollection := quickstartDatabase.Collection("Post")

	ping(client, ctx)
	var document interface{}

	document = bson.D{
		{"ID", 1},
		{"Name", "Kanit Mann"},
		{"Age", 19},
	}
	insertOneResult, err := insertOne(client, ctx, "Key",
		"Value", document)
	if err != nil {
		panic(err)
	}
	fmt.Println("Result of InsertOne")
	fmt.Println(insertOneResult.InsertedID)

	UserID, err := TaskCollection.InsertOne(ctx, bson.D{
		{Key: "PostID", Value: "OpenSourceIntro"},
		{Key: "Author", Value: "Kanit Mann"},
		{Key: "tags", Value: bson.A{"development", "programming", "coding"}},
	})

	postCollection, err := episodesCollection.InsertMany(ctx, []interface{}{
		bson.D{
			{"User ID", UserID.InsertedID},
			{"Name", "Kanit Mann"},
			{"Age", 20},
		},
		bson.D{
			{"User ID", UserID.InsertedID},
			{"Name", "Gaurav"},
			{"Age", 27},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted %v documents into Post collection!\n", len(postCollection.InsertedIDs))

}
