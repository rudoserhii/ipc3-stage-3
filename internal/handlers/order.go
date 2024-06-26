package handlers

import "github.com/gin-gonic/gin"

type OrderRoutes struct{}

func (p *OrderRoutes) GetAllOrders(*gin.Context) {
}

func (p *OrderRoutes) GetOrderByID(*gin.Context) {
}

func (p *OrderRoutes) CreateOrder(*gin.Context) {
}

func (p *OrderRoutes) CancelOrderByID(*gin.Context) {
}
