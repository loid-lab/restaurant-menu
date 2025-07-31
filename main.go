package main

import (
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/loid-lab/restaurant-menu/controllers"
	"github.com/loid-lab/restaurant-menu/initializers"
	"github.com/loid-lab/restaurant-menu/middleware"
	"github.com/loid-lab/restaurant-menu/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
	initializers.DB.AutoMigrate(
		&models.User{},
		&models.MenuItem{},
		&models.Order{},
		&models.Cart{},
		&models.CartItem{},
		&models.OrderItem{},
		&models.Payment{},
		&models.Invoice{},
		&models.InvoiceItem{},
	)

	initializers.ConnectCloudinary()
}

func main() {
	gin.SetMode(gin.DebugMode)

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Cors configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Split(os.Getenv("CORS_ORIGINS"), ","),
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	// Public Routes
	r.POST("/auth/register", controllers.CreateUser)
	r.POST("/auth/login", controllers.Login)
	r.POST("/webhook/stripe", controllers.StripeWebhook)

	// admin routes
	admin := r.Group("/admin")
	admin.Use(middleware.CheckAuth, middleware.CheckAdmin)
	{
		admin.GET("/metrics/sales", controllers.GetSaleMetrics)
		admin.GET("/orders/stats", controllers.GetOrderStats)
		admin.GET("/invoices", controllers.GetAllInvoices)
	}

	// authenticated user routes
	auth := r.Group("/user")
	auth.Use(middleware.CheckAuth)
	{
		auth.GET("/profile", controllers.GetUserProfile)

		auth.GET("/cart", controllers.GetCart)
		auth.POST("/cart/item", controllers.AddToCart)
		auth.POST("/cart/item/:id", controllers.DeleteCartItem)

		auth.POST("/order", controllers.CreateOrder)
		auth.GET("/orders", controllers.GetUserOrder)
		auth.POST("/orders/:id", controllers.GetOrderByID)

		auth.POST("/orders/:id/pay", controllers.CreateStripeCheckoutSession)

		auth.POST("menu", middleware.CheckAdmin, controllers.CreateMenuItem)
		auth.PUT("/menu/:id", middleware.CheckAdmin, controllers.UpdateMenuItems)
		auth.DELETE("/menu/:id", middleware.CheckAdmin, controllers.DeleteMenuItem)
	}

	// Public menu route
	r.GET("/menu", controllers.GetAllMenuItems)
	r.GET("/menu/:id", controllers.GetMenuItemsByID)

	r.Run()
}
