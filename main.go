package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	// Connect
	connStr := os.Getenv("CONNECTION_STRING")

	cxt, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	client, err := mongo.Connect(cxt, options.Client().ApplyURI(connStr))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(cxt)

	// Time how long ping takes
	ping := func() error {
		start := time.Now()
		err := client.Ping(cxt, readpref.Primary())
		stop := time.Now()
		fmt.Println("Ping", stop.Sub(start).String())
		return err
	}

	// Ping once
	ping()

	// Ping like you mean it
	err = ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected!")
}
