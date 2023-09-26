package router

import (
	"ecommerce/controller"

	"github.com/gin-gonic/gin"
)

// Router creates and configures the Gin router.
func Router() *gin.Engine {
	router := gin.Default()

	// Serve static files for specific routes
	router.Static("/index", "./frontend/index")
	router.Static("/home", "./frontend/home")
	router.Static("/signup", "./frontend/signup")
	router.Static("/signin", "./frontend/signin")
	router.Static("/additems", "./frontend/inventory")

	// Initialize session middleware

	// Define your routes
	router.GET("/getalldata", controller.Getalldata)
	router.POST("/create", controller.CreateProfile)
	router.POST("/addtocart", controller.Addtocart)
	router.POST("/login", controller.Login)
	router.POST("/products", controller.Products)
	router.POST("/updatecart", controller.UpdateCart)
	router.POST("/inventory", controller.Inventory)
	router.POST("/search", controller.Search)
	router.GET("/inventorydata", controller.Getallinventorydata)

	// Protected route that requires authentication
	return router
}
