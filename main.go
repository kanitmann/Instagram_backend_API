package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client
var MyMap map[string]*total_users

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

func find_post_id(w http.ResponseWriter, r *http.Request) {

	keys, ok := r.URL.Query()["ID"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	key := keys[0]

	collection := client.Database("Posts").Collection("Post")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	cursor, err := collection.Find(ctx, bson.M{"ID": key})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	var jsonDocuments []map[string]interface{}
	var bsonDocument bson.D
	var jsonDocument map[string]interface{}
	var temporaryBytes []byte
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {

		err = cursor.Decode(&bsonDocument)

		temporaryBytes, err = bson.MarshalExtJSON(bsonDocument, true, true)

		err = json.Unmarshal(temporaryBytes, &jsonDocument)

		jsonDocuments = append(jsonDocuments, jsonDocument)
		fmt.Println(jsonDocuments)
	}
	if err := cursor.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(w).Encode(jsonDocuments)

}

//Add Post

func addMeeting(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "invalid_http_method")
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))

	if err != nil {
		panic(err)
	}

	//Checking RACE Conditions
	res1 := strings.Split(string(body), "[{")
	res2 := strings.Split(res1[1], "}]")
	res3 := strings.ReplaceAll(res2[0], "} , {", ",")
	res5 := strings.Split(res2[1], ":")
	res6 := strings.Split(res5[1], ",")
	end_time, err := strconv.Atoi(strings.Trim(res5[2], "}"))
	start_time, err := strconv.Atoi(strings.Trim(res6[0], " "))
	res4 := strings.Split(res3, ",")
	for i, s := range res4 {
		a := strings.Split(s, ":")
		e1 := " "
		if (i+1)%2 == 0 {
			e1 = strings.Trim(a[1], "\"")
		}
		if (i+1)%3 == 0 {

			if a[1] == "\"YES\"" {
				obj := &total_users{[]int{}, []int{}}
				obj.AppendValues(start_time, end_time)
				MyMap[e1] = obj
				fmt.Fprintf(w, "ADDED WITHOUT COLLISION")
			}

		}
		fmt.Println(MyMap)
	}

	var m interface{}
	errr := bson.UnmarshalExtJSON([]byte(body), true, &m)
	if errr != nil {
		log.Println(errr)
	}
	log.Println(m)

	collection := client.Database("Posts").Collection("Post")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, m)

	json.NewEncoder(w).Encode(result)
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

	fmt.Println("Starting the application...")

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
			{"demo", "demo"},
		},
		bson.D{
			{"User ID", UserID.InsertedID},
			{"Name", "Gaurav"},
			{"Age", 27},
			{"demo", "demo"},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted %v documents into Post collection!\n", len(postCollection.InsertedIDs))

}
