package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	db *gorm.DB
)

// Models
type Seller struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Products []Product
}

type Product struct {
	gorm.Model
	SellerID   uint
	Name       string
	Description string
	Price      float64
	Quantity   uint
}

type Customer struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Cart     []CartItem
	Orders   []Order
}

type CartItem struct {
	gorm.Model
	CustomerID uint
	ProductID  uint
	Quantity   uint
}

type Order struct {
	gorm.Model
	CustomerID uint
	ProductID  uint
	Quantity   uint
	Status     string
}

// JWT authentication
func authMiddleware(c *gin.Context) {
	// Get the JWT token from the request header
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	c.Next()
}

// Seller authentication endpoints
func sellerLogin(c *gin.Context) {
	
}

func sellerRegister(c *gin.Context) {
	
}

// Seller product management endpoints
func addProduct(c *gin.Context) {
	
	// Parse the request body to get the product details
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create the product with the seller ID
	// product.SellerID = sellerID
	// db.Create(&product)

	c.JSON(http.StatusCreated, gin.H{"message": "Product created"})
}

func updateProduct(c *gin.Context) {
	
	// Get the product ID from the path parameter
	productID := c.Param("id")

	// Parse the request body to get the updated product details
	var updatedProduct Product
	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}



	c.JSON(http.StatusOK, gin.H{"message": "Product updated"})
}

func deleteProduct(c *gin.Context) {
	

	// Get the product ID from the path parameter
	productID := c.Param("id")



	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}

// Seller order management endpoints
func acceptOrder(c *gin.Context) {
	
	
	orderID := c.Param("id")

	

	c.JSON(http.StatusOK, gin.H{"message": "Order accepted"})
}

func rejectOrder(c *gin.Context) {
	

	
	orderID := c.Param("id")

	

	c.JSON(http.StatusOK, gin.H{"message": "Order rejected"})
}

// Customer product viewing and searching endpoints
func getAllProducts(c *gin.Context) { 

	
	

func searchProducts(c *gin.Context) {
	
	searchQuery := c.Query("q")


}


// Cart management endpoints
func addToCart(c *gin.Context) {


	
	var cartItem CartItem
	if err := c.ShouldBindJSON(&cartItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	
	c.JSON(http.StatusCreated, gin.H{"message": "Item added to cart"})
}

func updateCartItemQuantity(c *gin.Context) {
	// Get the customer ID from the context
	// customerID := c.GetInt("customerID")

	// Get the cart item ID from the path parameter
	cartItemID := c.Param("id")

	// Parse the request body to get the updated quantity
	var updatedQuantity struct {
		Quantity uint `json:"quantity"`
	}
	if err := c.ShouldBindJSON(&updatedQuantity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	

	c.JSON(http.StatusOK, gin.H{"message": "Cart item quantity updated"})
}

func removeFromCart(c *gin.Context) {
	// Get the customer ID from the context
	// customerID := c.GetInt("customerID")

	// Get the cart item ID from the path parameter
	cartItemID := c.Param("id")

	

	c.JSON(http.StatusOK, gin.H{"message": "Item removed from cart"})
}

// Checkout and order placement endpoints
func checkout(c *gin.Context) {
	// Get the customer ID from the context
	// customerID := c.GetInt("customerID")

	// Parse the request body to get the shipping address and payment information
	var checkoutData struct {
		ShippingAddress string `json:"shipping_address"`
		PaymentInfo     string `json:"payment_info"`
	}
	if err := c.ShouldBindJSON(&checkoutData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	c.JSON(http.StatusOK, gin.H{"message": "Order placed"})
}

func main() {
	// Connect to the database
	var err error
	db, err = gorm.Open("sqlite3", "database.db")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Migrate the database schema
	db.AutoMigrate(&Seller{}, &Product{}, &Customer{}, &CartItem{}, &Order{})

	// Create a new Gin router
	router := gin.Default()

	// Seller authentication routes
	router.POST("/seller/login", sellerLogin)
	router.POST("/seller/register", sellerRegister)

	// Protected routes (requires authentication)
	auth := router.Group("/api")
	auth.Use(authMiddleware)

	// Seller product management routes
	auth.POST("/products", addProduct)
	auth.PUT("/products/:id", updateProduct)
	auth.DELETE("/products/:id", deleteProduct)

	// Seller order management routes
	auth.PUT("/orders/:id/accept", acceptOrder)
	auth.PUT("/orders/:id/reject", rejectOrder)

	// Customer product viewing and searching routes
	router.GET("/products", getAllProducts)
	router.GET("/products/search", searchProducts)

	// Customer cart management routes
	auth.POST("/cart", addToCart)
	auth.PUT("/cart/:id", updateCartItemQuantity)
	auth.DELETE("/cart/:id", removeFromCart)

	// Customer checkout and order placement route
	auth.POST("/checkout", checkout)

	// Start the server
	fmt.Println("Server running on http://localhost:8080")
	router.Run(":8080")
}