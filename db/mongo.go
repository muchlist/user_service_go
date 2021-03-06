package db

import (
	"context"
	"github.com/muchlist/erru_utils_go/logger"
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
	// Db objek sebagai database objek
	Db *mongo.Database

	mongoURL = os.Getenv(mongoURLGetKey)
)

//Init menginisiasi database
//kembaliannya digunakan untuk memutus koneksi apabila main program dihentikan
func Init() (*mongo.Client, context.Context, context.CancelFunc) {

	Client, err := mongo.NewClient(options.Client().ApplyURI(mongoURL))
	if err != nil {
		logger.Error("Gagal membuat client Mongodb", err)
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)

	err = Client.Connect(ctx)
	if err != nil {
		logger.Error("Gagal menghubungkan MongoDb", err)
		panic(err)
	}

	Db = Client.Database("user_go")

	// err = Client.Ping(ctx, readpref.Primary())
	// if err != nil {
	// 	panic(err)
	// }

	logger.Info("MongoDb berhasil terkoneksi!")

	return Client, ctx, cancel
}
