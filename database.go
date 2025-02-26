package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var DB *mongo.Client

func ConnectDB() {
	uri := "mongodb://localhost:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Println("Error in creating MongoDB client", err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	// err = client.Connect(ctx)
	// cancel()
	if err != nil {
		fmt.Println("Error in connecting MongoDB", err)
	}
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		fmt.Println("Error in pinging MongoDB:", err)
		return
	}
	fmt.Println("Connected to MongoDB")
	DB = client
}
