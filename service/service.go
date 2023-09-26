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
	filter := bson.D{
		{"$and", []interface{}{
			bson.D{{"customerid", cart.CustomerId}},
			bson.D{{"name", cart.Name}},
		}},
	}
	_, err := config.Cart_Collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	update := bson.M{"$set": bson.M{"name": cart.Name, "quantity": cart.Quantity, "totalprice": cart.Price, "price": cart.Price / float64(cart.Quantity)}}
	result, err := config.Cart_Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	return result

}

func Cart(customerid string) []models.Cart {
	fmt.Println("1")
	filter :=  bson.M{"customerid":customerid}
	cursor, err := config.Cart_Collection.Find(context.Background(), filter)
	if err != nil {
		fmt.Println("2")
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	var Cart []models.Cart
	for cursor.Next(context.Background()) {
		var cart models.Cart
		err := cursor.Decode(&cart)
		if err != nil {
			fmt.Println("3")
			log.Fatal(err)
		}
		Cart = append(Cart, cart)
	}
	return Cart
}
func Search(productName string) []models.Inventory {
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
	filter := bson.M{"email": profile.Email}
	cursor, err := config.Customer_Collection.Find(context.Background(), filter)
	defer cursor.Close(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	profile.CustomerId = GenerateUniqueCustomerID()

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

func Addtocart(addtocart models.Addtocart1) bool {
	filter := bson.D{
		{"$and", []interface{}{
			bson.D{{"customerid", addtocart.CustomerId}},
			bson.D{{"name", addtocart.Name}},
		}},
	}

	cursor, err := config.Cart_Collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.Background())
	type addcart struct {
		CustomerId string  `json:"customerid" bson:"customerid"`
		Name       string  `json:"name" bson:"name"`
		Price      float64 `json:"price" bson:"price"`
		Quantity   int32   `json:"quantity" bson:"quantity"`
	}
	// Check if any items were found
	if cursor.RemainingBatchLength() == 0 {
		// Item not found, so insert a new item with quantity 1
		cart := addcart{CustomerId:addtocart.CustomerId, Name: addtocart.Name, Price: addtocart.Price, Quantity: 1}
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

func Login(details models.Login) (string, bool, error) {
	var customer models.Customer
	
	filter := bson.M{"email": details.Email}
	err := config.Customer_Collection.FindOne(context.Background(), filter).Decode(&customer)
	if err != nil {
		// Handle the case where the user is not found
		fmt.Println("1")
		return "", false, err
	}

	// Verify the password (You should use a secure password hashing library here)
	if customer.Password != details.Password {
		fmt.Println("2")
		// Passwords don't match
		return "", false, nil
	}

	token, err := CreateToken(customer.Email, customer.CustomerId)
	if err != nil {
		fmt.Println("3")
		return "", false, err

	}

	return token, true, nil
}
func Inventory(inventory models.Inventory) int {
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
