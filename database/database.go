package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoCtx context.Context
var client *mongo.Client
var database *mongo.Database

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' env")
	}

	client, err := mongo.Connect(mongoCtx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	database = client.Database("klever")
	fmt.Println("Connected to database")
}

func Stop() {
	if err := client.Disconnect(mongoCtx); err != nil {
		panic(err)
	}
	fmt.Println("Connection to database closed")
}
