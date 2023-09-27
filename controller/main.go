package controller

import (
	"ecommerce/constants"
	"ecommerce/models"
	"ecommerce/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Getinventorydata(c *gin.Context) {
	Inventorydata := service.Getinventorydata()
	fmt.Println(Inventorydata)
	c.JSON(http.StatusOK, Inventorydata)
}
func Getalldata(c *gin.Context) {
	alltransaction := service.Getalldata()
	fmt.Println(alltransaction)
	c.JSON(http.StatusOK, alltransaction)
}
func CreateSeller(c *gin.Context) {
	var seller models.Seller
	if err := c.BindJSON(&seller); err != nil {
		fmt.Println("error")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	fmt.Println(seller)
	createseller := service.CreateSeller(seller)

	c.JSON(http.StatusOK, createseller)

}
func Getallsellerdata(c *gin.Context) {
	fmt.Println("In seller")
	Getallsellerdata := service.Getallsellerdata()
	fmt.Println(Getallsellerdata)
	c.JSON(http.StatusOK, Getallsellerdata)

}
func UpdateCart(c *gin.Context) {
	var cart models.Cart
	if err := c.BindJSON(&cart); err != nil {
		fmt.Println("error")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	t1, err := service.ExtractCustomerID(cart.CustomerId, constants.SecretKey)
	if err != nil {
		fmt.Println("errorin Extaxt Token")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Token"})
		return
	}
	cart.CustomerId = t1

	result := service.UpdateCart(cart)
	c.JSON(http.StatusOK, result)
}
func Login(c *gin.Context) {
	var request models.Login
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	token, success, err := service.Login(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	if success {
		c.JSON(http.StatusOK, gin.H{"token": token})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}

func Products(c *gin.Context) {
	var cartproducts *models.Addtocart

	if err := c.BindJSON(&cartproducts); err != nil {
		fmt.Println("error")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	token, err := service.ExtractCustomerID(cartproducts.Token, constants.SecretKey)
	if err != nil {
		fmt.Println("errorin Extaxt Token")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Token"})
		return
	}
	cart := service.Cart(token)
	fmt.Println(cart)
	c.JSON(http.StatusOK, cart)
}

func Addtocart(c *gin.Context) {
	var addtocart models.Addtocart
	if err := c.BindJSON(&addtocart); err != nil {
		fmt.Println("error")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	fmt.Println(addtocart.Token)
	fmt.Println(constants.SecretKey)
	token, err := service.ExtractCustomerID(addtocart.Token, constants.SecretKey)
	if err != nil {
		fmt.Println("errorin Extaxt Token")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Token"})
		return
	}
	var addtocart1 models.Addtocart1
	addtocart1.CustomerId = token
	addtocart1.Name = addtocart.Name
	addtocart1.Price = addtocart.Price

	fmt.Println(addtocart)
	result := service.Addtocart(addtocart1)
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

func Update(c *gin.Context){
	var update models.Update
	if err := c.BindJSON(&update); err != nil {
		fmt.Println("error")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	fmt.Println(update)
	result := service.Update(update)
	c.JSON(http.StatusOK, result)
	
}

func Delete(c *gin.Context){
	var delete models.Delete
	if err := c.BindJSON(&delete); err != nil {
		fmt.Println("error")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	fmt.Println(delete)
	result := service.Delete(delete)
	fmt.Println(result)
	c.JSON(http.StatusOK, result)


}


func CheckSeller(c *gin.Context){
	var check models.Login
	if err := c.BindJSON(&check); err != nil {
		fmt.Println("error")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	fmt.Println(check)
	result := service.CheckSeller(check)
	fmt.Println(result)
	c.JSON(http.StatusOK, result)

}
