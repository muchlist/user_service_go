package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo/readpref"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	connectTimeout = 10
	mongoURLGetKey = "MONGO_DB_URL"
)

var (
	// Client sebagai satu sumber database client
	Client *mongo.Client

	mongoURL = os.Getenv(mongoURLGetKey)
)

func Init() {

	Client, err := mongo.NewClient(options.Client().ApplyURI(mongoURL))
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)
	defer cancel()

	err = Client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	err = Client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to MongoDB!")
}
