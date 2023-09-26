package config

import (
	"context"
	"fmt"
	"ecommerce/constants"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



var Customer_Collection *mongo.Collection
var Cart_Collection *mongo.Collection
var Inventory_Collection *mongo.Collection


func init() {
	clientoption := options.Client().ApplyURI(constants.Connectstring)

	client, err := mongo.Connect(context.TODO(), clientoption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDb sucessfully connected")
	Customer_Collection = client.Database(constants.DB_Name).Collection(constants.Customer_collection)
    Cart_Collection = client.Database(constants.DB_Name).Collection(constants.Cart_collection)
	Inventory_Collection = client.Database(constants.DB_Name).Collection(constants.Inventory_Collection)
	fmt.Println("Collection Connected")
}
