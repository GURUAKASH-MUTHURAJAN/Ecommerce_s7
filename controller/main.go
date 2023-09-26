package controller

import (
	"ecommerce/models"
	"ecommerce/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Getalldata(c *gin.Context) {
	alltransaction := service.Getalldata()
	c.JSON(http.StatusOK, alltransaction)
}

func UpdateCart(c *gin.Context) {
	var cart models.Cart
	if err := c.BindJSON(&cart); err != nil {
		fmt.Println("error")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	fmt.Println(cart)
	result := service.UpdateCart(cart)
	c.JSON(http.StatusOK, result)
}
func Login(c *gin.Context) {
	var Details models.Login
	if err := c.BindJSON(&Details); err != nil {
		fmt.Println("error")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	result := service.Login(Details)
	fmt.Println(result)
	if result {
		session := sessions.Default(c)
		session.Set("authenticated", true)
		if err := session.Save(); err != nil {
			fmt.Println("Error saving session:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}
		auth := session.Get("authenticated")
		fmt.Println("auth in controller",auth)
		fmt.Println("No error")
	}
	c.JSON(http.StatusOK, result)

}

func Products(c *gin.Context) {
	cart := service.Cart()
	c.JSON(http.StatusOK, cart)
}

func Addtocart(c *gin.Context) {
	var addtocart models.Addtocart
	if err := c.BindJSON(&addtocart); err != nil {
		fmt.Println("error")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	fmt.Println(addtocart)
	result := service.Addtocart(addtocart)
	c.JSON(http.StatusOK, result)

}

func CreateProfile(c *gin.Context) {
	var profile models.Customer
	if err := c.BindJSON(&profile); err != nil {
		fmt.Println("error")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	fmt.Println(profile)
	result := service.Insert(profile)
	c.JSON(http.StatusOK, result)
}
func Inventory(c *gin.Context) {
	var inventory models.Inventory
	fmt.Println("in inventory")
	if err := c.BindJSON(&inventory); err != nil {
		fmt.Println("error")
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	fmt.Println(inventory)
	result := service.Inventory(inventory)
	fmt.Println(result)
	c.JSON(http.StatusOK, result)

}
func Getallinventorydata(c *gin.Context) {
	result := service.Search(SearchName)
	fmt.Println(result)
	c.JSON(http.StatusOK, result)
}

var SearchName string

func Search(c *gin.Context) {
	type Serarch struct {
		ProductName string `json:"productName" bson:"productName"`
	}
	var search Serarch
	if err := c.BindJSON(&search); err != nil {
		fmt.Println("error")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	SearchName = search.ProductName
	fmt.Println(search)
}
