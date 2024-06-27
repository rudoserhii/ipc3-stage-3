package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/obiMadu/ipc3-stage-3/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (app *Config) routes() *gin.Engine {
	mux := gin.Default()

	mux.Use(cors.Default())

	// API version routes
	v1 := mux.Group("/api").Group("/v1")

	// Swagger docs
	v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Products
	products := v1.Group("/products")
	products.GET("/", app.Handlers.Product().GetAllProducts)
	products.GET("/:productID", app.Handlers.Product().GetProductByID)
	products.POST("/", app.Handlers.Product().CreateProduct)
	products.PUT("/:productID", app.Handlers.Product().UpdateProductByID)
	products.DELETE("/:productID", app.Handlers.Product().DeleteProductByID)

	// Product Images
	products.GET("/images/:imageName", app.Handlers.Product().GetProductImage)

	// Orders
	orders := v1.Group("/orders")
	orders.GET("/", app.Handlers.Order().GetAllOrders)
	orders.GET("/:orderID", app.Handlers.Order().GetOrderByID)
	orders.POST("/", app.Handlers.Order().CreateOrder)
	orders.DELETE("/:orderID", app.Handlers.Order().CancelOrderByID)

	return mux
}
