package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (app *Config) routes() *gin.Engine {
	mux := gin.Default()

	mux.Use(cors.Default())

	users := mux.Group("/api").Group("/v1").Group("/users")

	// User handlers
	// Products
	products := users.Group("/products")
	products.GET("/", app.Handlers.Product().GetAllProducts)
	products.GET("/:productID", app.Handlers.Product().GetProductByID)
	products.POST("/", app.Handlers.Product().CreateProduct)
	products.PUT("/:productID", app.Handlers.Product().UpdateProductByID)
	products.DELETE("/:productID", app.Handlers.Product().DeleteProductByID)

	// Orders
	orders := users.Group("/orders")
	orders.GET("/", app.Handlers.Order().GetAllOrders)
	orders.GET("/:orderID", app.Handlers.Order().GetOrderByID)
	orders.POST("/", app.Handlers.Order().CreateOrder)
	orders.DELETE("/:orderID", app.Handlers.Order().CancelOrderByID)

	return mux
}
