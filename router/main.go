package router

import (
	"ecommerce/controller"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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
	store := cookie.NewStore([]byte("your-secret-key"))
	router.Use(sessions.Sessions("mysession", store))

	// Define your routes
	router.GET("/getalldata", controller.Getalldata)
	router.POST("/create", controller.CreateProfile)
	router.POST("/addtocart", controller.Addtocart)
	router.POST("/login", controller.Login)
	router.GET("/products", controller.Products)
	router.POST("/updatecart", controller.UpdateCart)
	router.POST("/inventory", controller.Inventory)
	router.POST("/search", controller.Search)
	router.GET("/inventorydata", controller.Getallinventorydata)

	// Protected route that requires authentication
	router.GET("/home", AuthMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "This is a protected route."})
	})

	return router
}

// AuthMiddleware is a middleware function to check if a user is authenticated.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		auth := session.Get("authenticated")
		fmt.Println("auth in route",auth)
		fmt.Println("Session contents:", session)
		// Check if auth is not true or not set
		if auth != true {
			// Redirect to the login page or handle unauthorized access
			c.Redirect(http.StatusFound, "/signin")
			c.Abort()
			return
		}
		fmt.Println("Gone to next")
		c.Next()
	}
}

