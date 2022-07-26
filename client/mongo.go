package client

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/rjva-printerface/common/helpers"
	"github.com/rjva-printerface/common/helpers/colors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func getMongoUri() (mongoUri string) {
	uri := os.Getenv("MONGO_URI")

	if uri == "" {
		log.Fatal("MONGO_URI must be provided")
	}

	return uri
}

func CreateMongoClient() (mc *mongo.Client, cancel context.CancelFunc) {
	log := helpers.NewLog("mongo")
	mongoUri := getMongoUri()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+mongoUri))

	if err != nil {
		log.Print(err.Error(), colors.Red)
	}

	client.Database("book-service").Collection("books")

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Print(err.Error(), colors.Red)
	}

	return client, cancel
}
