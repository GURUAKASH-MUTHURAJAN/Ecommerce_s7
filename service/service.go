package service

import (
	"context"
	"ecommerce/config"
	"ecommerce/models"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateCart(cart models.Cart) *mongo.UpdateResult {
	filter := bson.M{"name": cart.Name}
	_, err := config.Cart_Collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("NAME", cart.Name)
	fmt.Println("Quantity", cart.Quantity)
	fmt.Println("price", cart.Price)
	fmt.Println("total", cart.TotalPrice)
	update := bson.M{"$set": bson.M{"name": cart.Name, "quantity": cart.Quantity, "totalprice": cart.Price, "price": cart.Price / float64(cart.Quantity)}}
	result, err := config.Cart_Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.UpsertedID)
	return result

}

func Cart() []models.Cart {
	filter := bson.D{}
	cursor, err := config.Cart_Collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	var Cart []models.Cart
	for cursor.Next(context.Background()) {
		var cart models.Cart
		err := cursor.Decode(&cart)
		if err != nil {
			log.Fatal(err)
		}
		Cart = append(Cart, cart)
	}
	return Cart
}
func Search(productName string) []models.Inventory{
	fmt.Println(productName)
	filter := bson.M{"itemcategory": productName}
	cursor, err := config.Inventory_Collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	var Inventory []models.Inventory
	for cursor.Next(context.Background()) {
		var inventory models.Inventory
		err := cursor.Decode(&inventory)
		if err != nil {
			log.Fatal(err)
		}
		Inventory = append(Inventory, inventory)
	}	
	fmt.Println(Inventory)
		return Inventory
}
func Getalldata() []primitive.M {
	filter := bson.D{}
	cursor, err := config.Customer_Collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	var Profiles []primitive.M
	for cursor.Next(context.Background()) {
		var profile bson.M
		err := cursor.Decode(&profile)
		if err != nil {
			log.Fatal(err)
		}
		Profiles = append(Profiles, profile)
	}
	return Profiles
}

func Insert(profile models.Customer) int {
	if profile.Password != profile.ConfirmPassword {
		return 3
	}
	filter := bson.M{"name": profile.Name}
	cursor, err := config.Customer_Collection.Find(context.Background(), filter)
	defer cursor.Close(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	if cursor.RemainingBatchLength() == 0 {
		fmt.Println("User name already exitst")
	inserted, err := config.Customer_Collection.InsertOne(context.Background(), profile)
	if err != nil {
		log.Fatal(err)
		return 0
	}
	fmt.Println("Inserted", inserted.InsertedID)
	return 1
}
return 2
}
func Addtocart(addtocart models.Addtocart) bool {
	fmt.Println(addtocart)
	filter := bson.M{"name": addtocart.Name}
	cursor, err := config.Cart_Collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("cursor")
	defer cursor.Close(context.Background())
	type addcart struct {
		Name     string  `json:"name" bson:"name"`
		Price    float64 `json:"price" bson:"price"`
		Quantity int32   `json:"quantity" bson:"quantity"`
	}
	// Check if any items were found
	if cursor.RemainingBatchLength() == 0 {
		// Item not found, so insert a new item with quantity 1
		cart := addcart{Name: addtocart.Name, Price: addtocart.Price, Quantity: 1}
		inserted, err := config.Cart_Collection.InsertOne(context.Background(), cart)
		if err != nil {
			log.Fatal(err)
			return false
		}
		fmt.Println("Inserted", inserted.InsertedID)
		return true
	}
	fmt.Println("Not in db")
	// Item already exists, update its quantity
	var cart addcart
	for cursor.Next(context.Background()) {
		err = cursor.Decode(&cart)
		fmt.Println(cart)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Decoded")
	}
	// Item already exists, update its quantity
	cart.Quantity++
	cart.Price = cart.Price + addtocart.Price
	// Use the UpdateOne method to increment the quantity
	update := bson.M{"$set": bson.M{"quantity": cart.Quantity, "price": cart.Price}}
	_, err = config.Cart_Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
		return false
	}
	fmt.Println("Updated quantity for", addtocart.Name)
	return true
}

func Login(details models.Login) bool {
	var customer models.Customer
	filter := bson.M{"name": details.Name}
	err := config.Customer_Collection.FindOne(context.Background(), filter).Decode(&customer)
	if err != nil {
		log.Fatal(err)
	}
	if customer.Password != details.Password {
		fmt.Println("Wrong password")
		return false
	}
	return true
}
func Inventory(inventory models.Inventory)int{
	//fmt.Println(inventory.Image)
	filter := bson.M{"itemname": inventory.ItemName}
	cursor, err := config.Inventory_Collection.Find(context.Background(), filter)
	defer cursor.Close(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	if cursor.RemainingBatchLength() == 0 {
	fmt.Println("No items available")
	inserted, err := config.Inventory_Collection.InsertOne(context.Background(), inventory)
	if err != nil {
		log.Fatal(err)
		return 0
	}
	fmt.Println("Inserted", inserted.InsertedID)
	return 1
    }
return 2

}
