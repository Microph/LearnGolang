package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	// Mongo URI pattern => mongodb://user:pass@host:port
	mongoURI := "mongodb://u30cbkfrc67odtpvtbyq:Zyl4Q2lK6zQKr4fRga4L@bnbfqre3af2qu3i-mongodb.services.clever-cloud.com:27017/bnbfqre3af2qu3i"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	collection := client.Database("bnbfqre3af2qu3i").Collection("tasks")
	ctx := context.TODO()
	cur, err := collection.Find(ctx, bson.D{})
	check(err)
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		check(err)

		fmt.Printf("%+v\n", result)
	}

	check(cur.Err())
}
