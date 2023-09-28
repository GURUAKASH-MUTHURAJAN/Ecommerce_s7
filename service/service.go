package service

import (
	"context"
	"ecommerce/config"
	"ecommerce/constants"
	"ecommerce/models"
	"fmt"
	"log"
	"regexp"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	filter := bson.M{"customerid": customerid}
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
func Search(productName string) []models.Inventory1 {
	fmt.Println(productName)
	filter := bson.M{"itemcategory": productName}
	cursor, err := config.Inventory_Collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	var Inventory []models.Inventory1
	for cursor.Next(context.Background()) {
		var inventory models.Inventory1
		err := cursor.Decode(&inventory)
		if err != nil {
			log.Fatal(err)
		}
		Inventory = append(Inventory, inventory)
	}
	return Inventory
}
func Getalldata() []models.Customer {
	filter := bson.D{}
	cursor, err := config.Customer_Collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	var Profiles []models.Customer
	for cursor.Next(context.Background()) {
		var profile models.Customer
		err := cursor.Decode(&profile)
		if err != nil {
			log.Fatal(err)
		}
		Profiles = append(Profiles, profile)
	}
	return Profiles
}
func Getinventorydata() []models.Inventory {
	filter := bson.D{}
	cursor, err := config.Inventory_Collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	var Inventorydata []models.Inventory
	for cursor.Next(context.Background()) {
		var inventory models.Inventory
		err := cursor.Decode(&inventory)
		if err != nil {
			log.Fatal(err)
		}
		Inventorydata = append(Inventorydata, inventory)
	}
	return Inventorydata
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
		cart := addcart{CustomerId: addtocart.CustomerId, Name: addtocart.Name, Price: addtocart.Price, Quantity: 1}
		_, err := config.Cart_Collection.InsertOne(context.Background(), cart)
		if err != nil {
			log.Fatal(err)
			return false
		}

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
	return true
}

func Getallsellerdata() []models.Seller {
	filter := bson.D{}
	cursor, err := config.Seller_Collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	var Seller []models.Seller
	for cursor.Next(context.Background()) {
		var seller models.Seller
		err := cursor.Decode(&seller)
		if err != nil {
			log.Fatal(err)
		}
		Seller = append(Seller, seller)
	}
	return Seller
}
func CreateSeller(seller models.Seller) bool {
	if seller.Password != seller.ConfirmPassword {
		return false
	}
	filter := bson.M{"selleremail": seller.Seller_Email}
	cursor, err := config.Seller_Collection.Find(context.Background(), filter)
	defer cursor.Close(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	if cursor.RemainingBatchLength() == 0 {
		seller.SellerId = GenerateUniqueCustomerID()
		_, err := config.Seller_Collection.InsertOne(context.Background(), seller)
		if err != nil {
			fmt.Println("error ")
			log.Fatal(err)

		}
		return true
	}
	return false
}
func Login(details models.Login) (string, bool, error) {
	var customer models.Customer

	filter := bson.M{"email": details.Email}
	err := config.Customer_Collection.FindOne(context.Background(), filter).Decode(&customer)
	if err != nil {
		// Handle the case where the user is not found
		return "", false, err
	}

	// Verify the password (You should use a secure password hashing library here)
	if customer.Password != details.Password {
		// Passwords don't match
		return "", false, nil
	}

	token, err := CreateToken(customer.Email, customer.CustomerId)
	if err != nil {
		return "", false, err

	}

	return token, true, nil
}
func Inventory(inventory models.Inventory) int {

	filter := bson.M{"itemname": inventory.ItemName}
	cursor, err := config.Inventory_Collection.Find(context.Background(), filter)
	defer cursor.Close(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	if cursor.RemainingBatchLength() == 0 {
		inventory1 := models.Inventory1{
			ItemCategory: inventory.ItemCategory,
			ItemName:     inventory.ItemName,
			Price:        inventory.Price,
			Quantity:     inventory.Quantity,
			Image:        inventory.Image,
		}
		var seller models.Seller
		err := config.Seller_Collection.FindOne(context.TODO(), bson.M{"sellerid": inventory.SellerId}).Decode(&seller)
		if err != nil {
			log.Fatal(err)
			return 0
		}
		inventory1.SellerName = seller.Seller_Name
		_, err1 := config.Inventory_Collection.InsertOne(context.Background(), inventory1)
		if err1 != nil {
			log.Fatal(err1)
			return 0
		}
		return 1
	}
	return 2

}

func Update(update models.Update) bool {
	if update.Collection == "seller" {
		filter := bson.M{"selleremail": update.IdName}
		update1 := bson.M{"$set": bson.M{update.Field: update.New_Value}}
		options := options.Update()
		_, err := config.Seller_Collection.UpdateOne(context.TODO(), filter, update1, options)
		if err != nil {
			fmt.Println("error while updating")
			return false
		}
		return true
	} else if update.Collection == "customer" {
		if update.Field == "phonenumber" || update.Field == "age" || update.Field == "pincode" {

			intValue, err := strconv.Atoi(update.New_Value)
			if err != nil {
				// Handle the error, e.g., return an error response or log it
			} else {
				update.New_Value = strconv.Itoa(intValue)
			}
			if !isValidNumber(update.New_Value) {
				return false
			}
			filter := bson.M{"email": update.IdName}
			update1 := bson.M{"$set": bson.M{update.Field: intValue}}
			options := options.Update()
			_, err1 := config.Customer_Collection.UpdateOne(context.TODO(), filter, update1, options)

			if err1 != nil {
				fmt.Println("error while updating")
				return false
			}

			return true
		}

		filter := bson.M{"email": update.IdName}
		update1 := bson.M{"$set": bson.M{update.Field: update.New_Value}}
		options := options.Update()
		_, err := config.Customer_Collection.UpdateOne(context.TODO(), filter, update1, options)
		if err != nil {
			fmt.Println("error while updating")
			return false
		}

		return true

	} else if update.Collection == "inventory" {
		if update.Field == "price" {
			// Check if New_Value is a valid integer
			intValue, err := strconv.Atoi(update.New_Value)
			if err != nil {
				// Handle the error, e.g., return an error response or log it
				return false
			}

			// Check if the input value is a valid number (numeric characters only)
			if !isValidNumber(update.New_Value) {
				return false
			}

			filter := bson.M{"itemname": update.IdName}
			update1 := bson.M{"$set": bson.M{update.Field: intValue}}
			options := options.Update()
			_, err1 := config.Inventory_Collection.UpdateOne(context.TODO(), filter, update1, options)
			if err1 != nil {
				fmt.Println("error while updating")
				return false
			}
			return true
		}

		filter := bson.M{"itemname": update.IdName}
		update1 := bson.M{"$set": bson.M{update.Field: update.New_Value}}
		options := options.Update()
		_, err := config.Inventory_Collection.UpdateOne(context.TODO(), filter, update1, options)
		if err != nil {
			fmt.Println("error while updating")
			return false
		}
		return true
	}

	return false
}

func Delete(delete models.Delete) bool {
	if delete.Collection == "customer" {
		filter := bson.M{"email": delete.IdValue}
		_, err := config.Customer_Collection.DeleteOne(context.Background(), filter)
		if err != nil {
			log.Fatal(err)
			return false
		}
		return true
	}
	if delete.Collection == "seller" {
		filter := bson.M{"selleremail": delete.IdValue}
		_, err := config.Seller_Collection.DeleteOne(context.Background(), filter)
		if err != nil {
			log.Fatal(err)
			return false
		}
		return true
	}
	if delete.Collection == "inventory" {
		filter := bson.M{"itemname": delete.IdValue}
		_, err := config.Inventory_Collection.DeleteOne(context.Background(), filter)
		if err != nil {
			log.Fatal(err)
			return false
		}
		return true
	}
	return true
}

func isValidNumber(s string) bool {
	numericRegex := regexp.MustCompile("^[0-9]+$")
	return numericRegex.MatchString(s)
}

func CheckSeller(check models.Login) (string, bool, error) {
	var seller models.Seller
	filter := bson.M{"selleremail": check.Email}
	config.Seller_Collection.FindOne(context.Background(), filter).Decode(&seller)
	if check.Password != seller.Password {
		return "", false, fmt.Errorf("InvalidPassword")
	}
	result, err := CreateToken(seller.Seller_Email, seller.SellerId)
	if err != nil {
		return "", false, err
	}
	fmt.Println(result)
	return result, true, nil
}

func DeleteProduct(delete models.DeleteProduct) bool {
	customerid, err := ExtractCustomerID(delete.Token, constants.SecretKey)
	if err != nil{
		log.Fatal(err)
		return false
	}
	filter1 := bson.M{"customerid": customerid}
	filter2 := bson.M{"name": delete.Name}
	combinedFilter := bson.M{
		"$and": []bson.M{filter1, filter2},
	}
	_, err = config.Cart_Collection.DeleteOne(context.Background(), combinedFilter)
	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}
