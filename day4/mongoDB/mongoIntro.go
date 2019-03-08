package main

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Mongo URI pattern => mongodb://user:pass@host:port
	mongoURI := "mongodb://u30cbkfrc67odtpvtbyq:Zyl4Q2lK6zQKr4fRga4L@bnbfqre3af2qu3i-mongodb.services.clever-cloud.com:27017/bnbfqre3af2qu3i"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
