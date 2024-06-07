package config

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database
var client *mongo.Client

func ConnectDB() {
    uri := os.Getenv("MONGODB_URI")
    clientOptions := options.Client().ApplyURI(uri)

    var err error
    client, err = mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    DB = client.Database("gofun_db")
    log.Println("Connected to MongoDB!")
}

func DisconnectDB() {
    if err := client.Disconnect(context.Background()); err != nil {
        log.Fatal(err)
    }
    log.Println("Disconnected from MongoDB!")
}
