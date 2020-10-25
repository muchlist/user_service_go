package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	connectTimeout = 10
	mongoURLGetKey = "MONGO_DB_URL"
)

var (
	// Client objek sebagai satu sumber database client
	Client *mongo.Client
	// Db objek sebagai database objek
	Db *mongo.Database

	mongoURL = os.Getenv(mongoURLGetKey)
)

//Init menginisiasi database
//kembaliannya digunakan untuk memutus koneksi apabila main program dihentikan
func Init() (*mongo.Client, context.Context, context.CancelFunc) {

	Client, err := mongo.NewClient(options.Client().ApplyURI(mongoURL))
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)

	err = Client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	Db = Client.Database("user_go")

	// err = Client.Ping(ctx, readpref.Primary())
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println("Connected to MongoDB!")

	return Client, ctx, cancel
}
