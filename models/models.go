package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Customer struct {
	CustomerId      string `json:"customerid" bson:"customerid"`
	Email           string `json:"email" bson:"email"`
	Name            string `json:"name" bson:"name"`
	Phone_No        int    `json:"phonenumber" bson:"phonenumber"`
	Age             int    `json:"age" bson:"age"`
	Password        string `json:"password" bson:"password"`
	ConfirmPassword string `json:"confirmpassword" bson:"confirmpassword"`
	FirstName       string `json:"firstname" bson:"firstname"`
	LastName        string `json:"lastname" bson:"lastname"`
	House_No        string `json:"houseno" bson:"houseno"`
	Street_Name     string `json:"streetname" bson:"streetname"`
	City            string `json:"city" bson:"city"`
	Pincode         int64  `json:"pincode" bson:"pincode"`
}
type Inventory struct {
	ItemCategory string  `json:"itemcategory" bson:"itemcategory"`
	ItemName     string  `json:"itemname" bson:"itemname"`
	Price        float64 `json:"price" bson:"price"`
	Quantity     string  `json:"quantity" bson:"quantity"`
	Image        string  `json:"image" bson:"image"`
}

type Addtocart struct {
	Token string  `json:"token" bson:"token"`
	Name  string  `json:"name" bson:"name"`
	Price float64 `json:"price" bson:"price"`
}
type Addtocart1 struct {
	CustomerId string  `json:"customerid" bson:"customerid"`
	Name  string  `json:"name" bson:"name"`
	Price float64 `json:"price" bson:"price"`
}
type Login struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}
type Cart struct {
	CustomerId string `json:"token"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
}
type Orders struct {
	Item_id    string  `json:"itemid" bson:"itemid"`
	Item_Name  string  `json:"itemname" bson:"itemname"`
	Quantity   int64   `json:"quantity" bson:"quantity"`
	Total_Cost float64 `json:"totalcost" bson:"totalcost"`
}
type Customer_Response struct {
	Id              primitive.ObjectID `json:"_id" bson:"_id"`
	Name            string             `json:"name" bson:"name"`
	Phone_No        int                `json:"phonenumber" bson:"phonenumber"`
	Age             int                `json:"age" bson:"age"`
	Password        string             `json:"password" bson:"password"`
	ConfirmPassword string             `json:"confirmpassword" bson:"confirmpassword"`
	FirstName       string             `json:"firstname" bson:"firstname"`
	LastName        string             `json:"lastname" bson:"lastname"`
	House_No        string             `json:"houseno" bson:"houseno"`
	Street_Name     string             `json:"streetname" bson:"streetname"`
	City            string             `json:"city" bson:"city"`
	Pincode         int64              `json:"pincode" bson:"pincode"`
}
