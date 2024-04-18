package database

import (
	"context"
	// "os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func dbinstance() *mongo.Client {
	// dsn := os.Getenv("MONGO_URL")
	mongodb_uri := "mongodb+srv://femi:Omojesu1992@cluster0.gmml5r3.mongodb.net/"
	// if dsn == "" {
    //     panic("MONGODB_URI environment variable is not set")
    // }
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongodb_uri))
	if err != nil {
		panic(err)
	}
	return client
}

var Client *mongo.Client = dbinstance()

var BookCollection *mongo.Collection = Client.Database("BookStoreManagement").Collection("BookCollection")
var UserCollection *mongo.Collection = Client.Database("BookStoreManagement").Collection("UserCollection")
var OrderCollection *mongo.Collection = Client.Database("BookStoreManagement").Collection("OrderCollection")
