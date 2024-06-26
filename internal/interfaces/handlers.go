package interfaces

import "github.com/gin-gonic/gin"

type Handlers interface {
	Product() ProductRoutes
	Order() OrderRoutes
}

type ProductRoutes interface {
	GetAllProducts(c *gin.Context)
	GetProductByID(c *gin.Context)
	CreateProduct(c *gin.Context)
	UpdateProductByID(c *gin.Context)
	DeleteProductByID(c *gin.Context)
}

type OrderRoutes interface {
	GetAllOrders(c *gin.Context)
	GetOrderByID(c *gin.Context)
	CreateOrder(c *gin.Context)
	CancelOrderByID(c *gin.Context)
}
