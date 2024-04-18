package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func dbinstance() *mongo.Client {
	mongodb_uri := "mongodb+srv://femi:Omojesu1992@cluster0.gmml5r3.mongodb.net/"
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
